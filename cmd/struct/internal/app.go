package internal

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xconvert"
	"github.com/evercyan/brick/xfile"
	"github.com/evercyan/brick/xutil"
	"github.com/xwb1989/sqlparser"
)

type (
	// App ...
	App struct{}
	// CommentInfo ...
	CommentInfo struct {
		Field string
		Type  string
		Title string
		Enums []map[string]string
	}
)

var (
	app = new(App)
)

// ----------------------------------------------------------------

// GormStruct ...
func (t *App) GormStruct(text string) (string, error) {
	return t.buildStruct(text, sceneGorm)
}

// CommonStruct ...
func (t *App) CommonStruct(text string) (string, error) {
	return t.buildStruct(text, sceneCommon)
}

// Sql ...
func (t *App) Sql(filepath, dstDir string) {
	// 读取 sql 文件内容
	content := xfile.Read(filepath)

	// 匹配 `CREATE TABLE 建表语句`
	re := regexp.MustCompile(`(?s)CREATE TABLE (.*?) \(.*?COMMENT='[^']*?';`)
	matches := re.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		xcolor.Fail("Error:", "无效的 sql 内容, 建议从数据库相关工具直接导出结构来使用")
		return
	}

	count := 0
	for _, match := range matches {
		dstFile := fmt.Sprintf("%s/%s.go", dstDir, strings.ReplaceAll(match[1], "`", ""))
		if xfile.IsExist(dstFile) {
			continue
		}

		sql := match[0]
		res, err := t.buildStruct(sql, sceneFile)
		if err != nil {
			xcolor.Fail("Error:", fmt.Sprintf("解析 sql 失败: %s \n %s", err.Error(), sql))
			continue
		}

		if err := xfile.Write(dstFile, res); err != nil {
			xcolor.Fail("Error:", fmt.Sprintf("写入文件失败: %s", err.Error()))
			continue
		}
		xcolor.Success("Success:", fmt.Sprintf("写入文件成功: %s", dstFile))
		count++
	}
	xcolor.Success("\n🍺🍺🍺", fmt.Sprintf("共生成 %d 个文件", count))
}

// Enum ...
func (t *App) Enum(text string) (string, error) {
	info := t.parseComment(text)
	if info == nil {
		return "", fmt.Errorf("invalid enum text")
	}
	return t.buildEnum(info), nil
}

// ----------------------------------------------------------------

// parseComment ...
func (t *App) parseComment(comment string) *CommentInfo {
	v := strings.Split(comment, ":")
	if len(v) != 4 {
		return nil
	}
	vv := strings.Split(v[3], ";")
	if len(vv) < 2 {
		return nil
	}
	res := &CommentInfo{
		Field: strings.TrimSpace(v[0]),
		Type:  strings.TrimSpace(v[1]),
		Title: strings.TrimSpace(v[2]),
		Enums: make([]map[string]string, 0),
	}
	index := 0
	for _, vvv := range vv {
		vvvv := strings.Split(vvv, ",")
		if len(vvvv) != 2 {
			continue
		}
		res.Enums = append(res.Enums, map[string]string{
			"index": fmt.Sprint(index),
			"val":   strings.TrimSpace(vvvv[0]),
			"title": strings.TrimSpace(vvvv[1]),
		})
		index++
	}
	return res
}

// buildEnum ...
func (t *App) buildEnum(info *CommentInfo) string {
	builder := strings.Builder{}
	builder.WriteString("\n")
	// type
	builder.WriteString(fmt.Sprintf("// %s %s\n", info.Field, info.Title))
	builder.WriteString(fmt.Sprintf("type %s %s\n\n", info.Field, info.Type))
	// const
	builder.WriteString("const (\n")
	for k, v := range info.Enums {
		if k == 0 {
			if v["val"] == "0" {
				builder.WriteString(fmt.Sprintf("\t%s%s %s = iota\n", info.Field, v["index"], info.Field))
			} else {
				builder.WriteString(fmt.Sprintf(
					"\t%s%s %s = iota + %s\n", info.Field, v["index"], info.Field, v["val"],
				))
			}
		} else {
			builder.WriteString(fmt.Sprintf("\t%s%s\n", info.Field, v["index"]))
		}
	}
	builder.WriteString(")\n\n")
	// func
	builder.WriteString(fmt.Sprintf("func (t %s) String() string {\n", info.Field))
	builder.WriteString("\tswitch t {\n")
	for _, v := range info.Enums {
		builder.WriteString(fmt.Sprintf("\tcase %s%s:\n", info.Field, v["index"]))
		builder.WriteString(fmt.Sprintf("\t\treturn \"%s\"\n", v["title"]))
	}
	builder.WriteString("\tdefault:\n")
	builder.WriteString("\t\treturn \"\"\n")
	builder.WriteString("\t}\n")
	builder.WriteString("}\n")
	return builder.String()
}

// buildStruct ...
func (t *App) buildStruct(text string, scene int) (string, error) {
	text = t.replaceKeyword(text, scene)
	statement, err := sqlparser.ParseStrictDDL(text)
	if err != nil {
		return "", err
	}
	stmt, ok := statement.(*sqlparser.DDL)
	if !ok || stmt.Action != sqlparser.CreateStr {
		return "", fmt.Errorf("sql is not a create statement")
	}
	// 构造输出
	builder := strings.Builder{}
	// 表名 & 注释
	tableName := stmt.NewName.Name.String()
	structName := strings.Title(xconvert.ToCamelCase(tableName))
	if scene == sceneFile {
		builder.WriteString("package po\n")
		builder.WriteString("\n")
		// 引入 time
		for _, column := range stmt.TableSpec.Columns {
			toType, ok := typeMap[column.Type.Type]
			if ok && toType == "time.Time" {
				builder.WriteString("import (\n")
				builder.WriteString("\t\"time\"\n")
				builder.WriteString(")\n\n")
				break
			}
		}
	} else {
		builder.WriteString("\n")
	}
	builder.WriteString(fmt.Sprintf("// %s ...\n", structName))
	builder.WriteString(fmt.Sprintf("type %s struct { \n", structName))
	// 枚举字段
	commentList := make([]*CommentInfo, 0)
	// 字段
	for _, column := range stmt.TableSpec.Columns {
		// 类型
		columnType := column.Type.Type
		if column.Type.Unsigned {
			columnType += " unsigned"
		}
		toType, ok := typeMap[columnType]
		if !ok {
			toType = "string"
		}
		// 字段原名
		field := column.Name.String()

		// 特殊字段处理
		if xutil.InArray(toType, typeNumber) {
			if strings.HasSuffix(field, "id") {
				toType = "uint64"
			} else if strings.HasSuffix(field, "time") {
				toType = "int64"
			}
		}

		// 字段大驼峰
		toField := strings.Title(xconvert.ToCamelCase(field))
		// JSON 可选择性使用下划线
		JSONField := xconvert.ToCamelCase(field)
		if FlagJSONUseSnake {
			JSONField = field
		}
		if scene == sceneCommon {
			// 标准的结构体
			builder.WriteString(fmt.Sprintf(
				"\t%s\t%s\t`json:\"%s\"`\n",
				toField,
				toType,
				JSONField,
			))
		} else {
			// gorm 结构体, 带枚举字段
			toComment := ""
			comment := column.Type.Comment
			if comment != nil {
				toComment = string(comment.Val)
			}
			if toComment != "" {
				commentInfo := t.parseComment(fmt.Sprintf("%s%s:%s:%s", structName, toField, toType, toComment))
				if commentInfo != nil {
					commentList = append(commentList, commentInfo)
					// 能解析出枚举, 则变更字段类型
					toType = structName + toField
				}
			}
			if FlagComment {
				builder.WriteString(fmt.Sprintf(
					"\t%s\t%s\t`json:\"%s\" gorm:\"column:%s\" comment:\"%s\"`\n",
					toField,
					toType,
					JSONField,
					field,
					toComment,
				))
			} else {
				builder.WriteString(fmt.Sprintf(
					"\t%s\t%s\t`json:\"%s\" gorm:\"column:%s\"`\n",
					toField,
					toType,
					JSONField,
					field,
				))
			}
		}
	}
	builder.WriteString("}\n")

	if scene == sceneGorm || scene == sceneFile {
		builder.WriteString("\n")
		builder.WriteString("// TableName ...\n")
		builder.WriteString(fmt.Sprintf("func (t *%s) TableName() string {\n", structName))
		builder.WriteString(fmt.Sprintf("\treturn \"%s\"\n", tableName))
		builder.WriteString("}\n")
	}

	for _, v := range commentList {
		builder.WriteString(t.buildEnum(v))
	}

	res := t.fmt(t.restoreKeyword(builder.String()))
	return res, nil
}

// ----------------------------------------------------------------

// replaceKeyword 替换关键字
func (t *App) replaceKeyword(s string, scene int) string {
	for _, keyword := range protectedFields {
		lastChar := keyword[len(keyword)-1:]
		if scene == sceneFile {
			s = strings.ReplaceAll(
				s,
				fmt.Sprintf("`%s`", keyword),
				fmt.Sprintf("`%s%s`", keyword, lastChar),
			)
		} else {
			s = strings.ReplaceAll(
				s,
				fmt.Sprintf(" %s ", keyword),
				fmt.Sprintf(" %s%s ", keyword, lastChar),
			)
		}
	}
	return s
}

// restoreKeyword 还原关键字
func (t *App) restoreKeyword(s string) string {
	for _, keyword := range protectedFields {
		lastChar := keyword[len(keyword)-1:]
		field := xconvert.ToCamelCase(keyword)
		title := strings.Title(field)
		// CheckCodee	string	`json:"checkCodee" gorm:"column:check_codee"`
		// CheckCodee => CheckCode
		// checkCodee => checkCode
		// check_codee => check_code
		s = xutil.Replace(s, map[string]string{
			title + lastChar:   title,
			field + lastChar:   field,
			keyword + lastChar: keyword,
		})
	}
	return s
}

// fmt gofmt 格式化
func (t *App) fmt(s string) string {
	filepath := os.TempDir() + "tmp.go"
	xfile.Write(filepath, s)
	output, err := exec.Command("gofmt", filepath).Output()
	if err != nil {
		return s
	}
	return string(output)
}

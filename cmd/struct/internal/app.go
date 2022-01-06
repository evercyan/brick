package internal

import (
	"fmt"
	"strings"

	"github.com/evercyan/brick/xconvert"
	"github.com/xwb1989/sqlparser"
)

type (
	App         struct{}
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

// GormStruct generate gorm struct
func (t *App) GormStruct(text string) (string, error) {
	return t.buildStruct(text, SceneGorm)
}

// Enum generate enum code
func (t *App) Enum(text string) (string, error) {
	info := t.parseComment(text)
	if info == nil {
		return "", fmt.Errorf("invalid enum text")
	}
	return t.buildEnum(info), nil
}

// CommonStruct generate common struct
func (t *App) CommonStruct(text string) (string, error) {
	return t.buildStruct(text, SceneCommon)
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
	statement, err := sqlparser.ParseStrictDDL(text)
	if err != nil {
		return "", err
	}
	stmt, ok := statement.(*sqlparser.DDL)
	if !ok || stmt.Action != sqlparser.CreateStr {
		return "", fmt.Errorf("sql is not a create statment")
	}
	// 构造输出
	builder := strings.Builder{}
	// 表名 & 注释
	tableName := stmt.NewName.Name.String()
	structName := strings.Title(xconvert.ToCamelCase(tableName))
	builder.WriteString("\n")
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
		// 字段大驼峰
		toField := strings.Title(xconvert.ToCamelCase(field))
		// JSON 可选择性使用下划线
		JSONField := xconvert.ToCamelCase(field)
		if FlagJSONUseSnake {
			JSONField = field
		}
		if scene == SceneCommon {
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
				commentInfo := t.parseComment(fmt.Sprintf("%s:%s:%s", toField, toType, toComment))
				if commentInfo != nil {
					commentList = append(commentList, commentInfo)
					// 能解析出枚举, 则变更字段类型
					toType = toField
				}
			}
			builder.WriteString(fmt.Sprintf(
				"\t%s\t%s\t`json:\"%s\" gorm:\"column:%s\" comment:\"%s\"`\n",
				toField,
				toType,
				JSONField,
				field,
				toComment,
			))
		}
	}
	builder.WriteString("}\n")

	if scene == SceneGorm {
		builder.WriteString("\n")
		builder.WriteString("// TableName ...\n")
		builder.WriteString(fmt.Sprintf("func (t *%s) TableName() string {\n", structName))
		builder.WriteString(fmt.Sprintf("\treturn \"%s\"\n", tableName))
		builder.WriteString("}\n")
	}

	for _, v := range commentList {
		builder.WriteString(t.buildEnum(v))
	}

	return builder.String(), nil
}

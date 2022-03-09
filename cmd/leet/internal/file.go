package internal

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strings"

	"github.com/evercyan/brick/cmd/leet/config"
	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xfile"
)

type File struct{}

func newFile() *File {
	return &File{}
}

// ----------------------------------------------------------------

// GenerateQuestion ...
func (t *File) GenerateQuestion(detail *config.Question, fileDir, lang string) error {
	// 答题目录
	questionDir := fileDir + GetQuestionPath(detail.Fid, detail.Qid, detail.Slug)
	if !xfile.IsExist(questionDir) {
		if err := os.MkdirAll(questionDir, os.ModePerm); err != nil {
			return fmt.Errorf("创建答题文件目录失败")
		}
	}

	// 生成答题 README.md
	tagList := make([]string, 0)
	for _, v := range detail.Detail.TagList {
		if v["slug"] == "" {
			continue
		}
		tagList = append(tagList, fmt.Sprintf(
			"[%s](%s)",
			v["name"],
			fmt.Sprintf(config.LeetCodeTagURL, v["slug"]),
		))
	}
	readmeInfo := struct {
		Id      string
		Title   interface{}
		Link    string
		Content interface{}
		Level   interface{}
		Tag     interface{}
	}{
		Id:      detail.Fid,
		Title:   template.HTML(detail.Title),
		Link:    detail.Link,
		Content: template.HTML(detail.Detail.Content),
		Level:   template.HTML(detail.Level.String()),
		Tag:     template.HTML(strings.Join(tagList, " ")),
	}
	var b bytes.Buffer
	tpl := template.Must(template.New("readme").Parse(config.TplQuestionReadme))
	if err := tpl.Execute(&b, readmeInfo); err != nil {
		fmt.Println(err)
		return fmt.Errorf("解析答题 README 模板失败")
	}
	fileReadme := fmt.Sprintf("%s/README.md", questionDir)
	if err := xfile.Write(fileReadme, string(b.Bytes())); err != nil {
		return fmt.Errorf("创建答题 README 文件失败")
	}

	// 生成答题文件和测试文件
	file, fileTpl, testfile, testfileTpl := "", "", "", ""
	if langMap, ok := config.LangMap[lang]; ok {
		file, fileTpl = langMap["file"], langMap["fileTpl"]
		testfile, testfileTpl = langMap["testfile"], langMap["testfileTpl"]
	} else if langExt, ok := config.LangExt[lang]; ok {
		file, fileTpl = fmt.Sprintf(config.LangFile, langExt), config.LangFileTpl
	} else {
		// 未配置编程语言的, 生成完 README.md 直接结束
		return nil
	}

	questionFile := fmt.Sprintf("%s/%s", questionDir, file)
	if xfile.IsExist(questionFile) {
		xcolor.Success(config.SymbolSuccess, fmt.Sprintf("答题文件已存在: %s", questionFile))
	} else {
		// javascript: 替换题目文件中 module.exports = FuncToReplace;
		if lang == "javascript" {
			matchs := regexp.MustCompile(`var ([^=]+) = function`).FindStringSubmatch(
				detail.Detail.CodeMap[lang],
			)
			if len(matchs) >= 2 {
				fileTpl = strings.Replace(fileTpl, "FuncToReplace", matchs[1], -1)
			}
		}
		err := xfile.Write(questionFile, fmt.Sprintf(fileTpl, detail.Detail.CodeMap[lang]))
		if err != nil {
			return fmt.Errorf("生成答题文件失败")
		}
	}

	if testfile != "" {
		questionTestFile := fmt.Sprintf("%s/%s", questionDir, testfile)
		if xfile.IsExist(questionTestFile) {
			xcolor.Success(config.SymbolSuccess, fmt.Sprintf("测试文件已存在: %s", questionTestFile))
		} else {
			// golang: 替换测试文件中函数名称
			if lang == "golang" {
				matchs := regexp.MustCompile(`func ([^(]+)\(`).FindStringSubmatch(
					detail.Detail.CodeMap[lang],
				)
				if len(matchs) >= 2 {
					testfileTpl = strings.Replace(testfileTpl, "FuncToReplace", matchs[1], -1)
				}
			}
			err := xfile.Write(questionTestFile, testfileTpl)
			if err != nil {
				return fmt.Errorf("生成测试文件失败")
			}
		}
	}

	return nil
}

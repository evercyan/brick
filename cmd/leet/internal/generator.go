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

// Generator ...
type Generator struct{}

func newGenerator() *Generator {
	return &Generator{}
}

// ----------------------------------------------------------------

// GenerateQuestion ...
func (t *Generator) GenerateQuestion(
	detail *config.Question,
	projectDir string,
	lang string,
) error {
	// 答题目录
	questionDir := fmt.Sprintf(
		"%s/%s", projectDir, GetQuestionPath(detail.Fid, detail.Qid, detail.Slug),
	)
	if !xfile.IsExist(questionDir) {
		if err := os.MkdirAll(questionDir, os.ModePerm); err != nil {
			return fmt.Errorf("创建答题文件目录失败")
		}
	}

	// 生成答题 README.md
	tagList := make([]string, 0)
	for _, v := range detail.Detail.TagList {
		if v.Slug == "" {
			continue
		}
		tagList = append(tagList, fmt.Sprintf(
			"[%s](%s)",
			v.Name,
			fmt.Sprintf(config.LeetCodeTagURL, v.Slug),
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

// GenerateRecord ...
func (t *Generator) GenerateRecord(
	list []*config.Question,
	projectDir string,
) error {
	lines := make([]string, 0)
	lines = append(lines, "| # | 标题 | 难度 |")
	lines = append(lines, "| :-: | :-- | :-: |")
	for _, v := range list {
		qPath := GetQuestionPath(v.Fid, v.Qid, v.Slug)
		qReadme := fmt.Sprintf("%s/%s/README.md", projectDir, qPath)
		if !xfile.IsExist(qReadme) {
			continue
		}
		lines = append(lines, fmt.Sprintf(
			"| [%s](%s) | [%s](%s) | %s |",
			v.Fid,
			v.Link,
			v.Title,
			fmt.Sprintf("./%s", qPath),
			v.Level.String(),
		))
	}
	recordInfo := struct {
		Question string
	}{
		Question: strings.Join(lines, "\n"),
	}
	var b bytes.Buffer
	tpl := template.Must(template.New("record").Parse(config.TplRecord))
	if err := tpl.Execute(&b, recordInfo); err != nil {
		return fmt.Errorf("解析答题纪录模板失败")
	}
	fileRecord := fmt.Sprintf("%s/RECORD.md", projectDir)
	if err := xfile.Write(fileRecord, string(b.Bytes())); err != nil {
		return fmt.Errorf("创建答题纪录文件失败")
	}
	return nil
}

// GenerateTag ...
func (t *Generator) GenerateTag(
	tagList []*config.Tag,
	tagMap map[string][]*config.Question,
	projectDir string,
) error {
	tagPath := fmt.Sprintf("%s/%s", projectDir, config.CategoryPath)
	if !xfile.IsExist(tagPath) {
		if err := os.MkdirAll(tagPath, os.ModePerm); err != nil {
			return fmt.Errorf("创建标签文件目录失败")
		}
	}
	tpl := template.Must(template.New("tag").Parse(config.TplTag))
	for _, tag := range tagList {
		questionList, ok := tagMap[tag.Slug]
		if !ok {
			continue
		}
		lines := make([]string, 0)
		lines = append(lines, "| # | 标题 | 难度 | 状态 |")
		lines = append(lines, "| :-: | :-- | :-: | :-: |")
		for _, v := range questionList {
			qPath := GetQuestionPath(v.Fid, v.Qid, v.Slug)
			qReadme := fmt.Sprintf("%s/%s/README.md", projectDir, qPath)
			status := ""
			if xfile.IsExist(qReadme) {
				status = "✅"
			}
			lines = append(lines, fmt.Sprintf(
				"| [%s](%s) | [%s](%s) | %s | %s |",
				v.Fid,
				v.Link,
				v.Title,
				fmt.Sprintf("../%s", qPath),
				v.Level.String(),
				status,
			))
		}
		badge := fmt.Sprintf(
			"![total](https://img.shields.io/badge/total-%d-ff9985.svg?style=flat)",
			tag.Count,
		)
		recordInfo := struct {
			Name     string
			Link     string
			Badge    string
			Question string
		}{
			Name:     tag.Name,
			Link:     GetTagLink(tag.Slug),
			Badge:    badge,
			Question: strings.Join(lines, "\n"),
		}
		var b bytes.Buffer
		if err := tpl.Execute(&b, recordInfo); err != nil {
			xcolor.Fail(config.SymbolError, "解析标签模板失败")
			continue
		}
		fileTag := fmt.Sprintf("%s/%s.md", tagPath, tag.Name)
		if err := xfile.Write(fileTag, string(b.Bytes())); err != nil {
			xcolor.Fail(config.SymbolError, "创建标签文件失败")
			continue
		}
		xcolor.Success(config.SymbolNotice, fileTag)
	}
	return nil
}

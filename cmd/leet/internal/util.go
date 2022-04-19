package internal

import (
	"fmt"
	"html"
	"os"
	"os/user"
	"regexp"
	"strings"

	"github.com/evercyan/brick/cmd/leet/config"
	"github.com/evercyan/brick/xconvert"
	"github.com/evercyan/brick/xfile"
)

// GetCfgPath ...
func GetCfgPath() string {
	userPath, err := user.Current()
	if err != nil {
		panic(err)
	}
	filepath := fmt.Sprintf(
		"%s/.config/%s",
		userPath.HomeDir,
		strings.ToLower(config.App),
	)
	if !xfile.IsExist(filepath) {
		if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
			panic(err)
		}
	}
	return filepath
}

// GetCfgFilepath ...
func GetCfgFilepath() string {
	return fmt.Sprintf("%s/%s", GetCfgPath(), config.ConfigFile)
}

// GetQuestionFilepath ...
func GetQuestionFilepath() string {
	return fmt.Sprintf("%s/%s", GetCfgPath(), config.QuestionFile)
}

// GetQuestionLink ...
func GetQuestionLink(slug string) string {
	return fmt.Sprintf(config.LeetCodeQuestionURL, slug)
}

// GetTagLink ...
func GetTagLink(tag string) string {
	return fmt.Sprintf(config.LeetCodeTagURL, tag)
}

// FormatQuestionContent ...
func FormatQuestionContent(s string) string {
	// 处理 img 标签, 生成 markdown 格式, ![](图片链接地址)
	rImg, _ := regexp.Compile(`<img[^<>]*src="([^"]+)"[^<>]*>`)
	s = rImg.ReplaceAllString(s, "```\n\n![]($1)\n\n```")
	// 处理 &nbsp; &lt; 等 html escape
	s = html.UnescapeString(s)
	// 处理多行换行
	s = regexp.MustCompile(`(\n){3,}`).ReplaceAllString(s, "\n\n")
	// 处理标签
	re, _ := regexp.Compile("<[^<>]+>")
	return re.ReplaceAllString(s, "")
}

// GetQuestionPath ...
func GetQuestionPath(fid string, qid int64, slug string) string {
	id := int64(xconvert.ToUint(fid))
	if id == 0 {
		id = qid
	}
	tpl := "%d-%s"
	if id < 10000 {
		tpl = "%04d-%s"
	}
	return fmt.Sprintf("%s/%s", config.QuestionPath, fmt.Sprintf(tpl, id, slug))
}

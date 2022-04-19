package internal

import (
	"fmt"
	"strings"

	"github.com/evercyan/brick/cmd/leet/config"
	"github.com/evercyan/brick/xcli/xtable"
	"github.com/evercyan/brick/xfile"
	"github.com/evercyan/brick/xutil"
	"gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Path         string `json:"path" yaml:"path" table:"答题目录"`
	Lang         string `json:"lang" yaml:"lang" table:"答题语言"`
	QuestionFile string `json:"question_file" yaml:"questionFile"`
}

func newConfig() *Config {
	res := &Config{}
	content := xfile.Read(GetCfgFilepath())
	if content == "" {
		return res
	}
	yaml.Unmarshal([]byte(content), res)
	return res
}

// ----------------------------------------------------------------

// Render ...
func (t *Config) Render() string {
	list := [][]string{
		{
			"答题目录", t.Path,
		},
		{
			"答题语言", t.Lang,
		},
	}
	return xtable.New(list).Style(xtable.Dashed).Border(true).Text()
}

// SetPath ...
func (t *Config) SetPath(s string) error {
	if !xfile.IsExist(s) {
		return fmt.Errorf("目录不存在, 请手动创建")
	}
	if !xfile.IsDir(s) {
		return fmt.Errorf("无效的文件目录")
	}
	t.Path = s
	return t.Update()
}

// SetLang ...
func (t *Config) SetLang(s string) error {
	lang := strings.ToLower(s)
	if !xutil.IsContains(config.LangList, lang) {
		return fmt.Errorf("无效的答题语言, 支持: [%s]", strings.Join(config.LangList, ", "))
	}
	t.Lang = lang
	return t.Update()
}

// Update ...
func (t *Config) Update() error {
	b, _ := yaml.Marshal(t)
	return xfile.Write(GetCfgFilepath(), string(b))
}

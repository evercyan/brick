package xfile

import (
	"os"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

// MinifyHTML ...
func MinifyHTML(srcPath, destPath string) error {
	m := minify.New()
	m.Add("text/html", &html.Minifier{
		KeepComments:            false, // 移除注释
		KeepConditionalComments: false, // 移除IE条件注释
		KeepDefaultAttrVals:     true,  // 保留默认属性值
		KeepDocumentTags:        true,  // 保留 html, head, body 标签
		KeepEndTags:             true,  // 保留所有结束标签
		KeepQuotes:              true,  // 尽可能移除属性值的引号
	})
	m.Add("text/css", &css.Minifier{})
	m.Add("application/javascript", &js.Minifier{
		KeepVarNames: false, // 保留变量名
	})
	input, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}
	output, err := m.Bytes("text/html", input)
	if err != nil {
		return err
	}
	return Write(destPath, string(output))
}

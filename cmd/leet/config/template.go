package config

import (
	_ "embed"
)

//go:embed question_readme.tpl
var TplQuestionReadme string

//go:embed question_go_test.tpl
var TplQuestionGoTest string

//go:embed question_js_test.tpl
var TplQuestionJsTest string

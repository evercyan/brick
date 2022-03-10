package config

import (
	_ "embed"
)

//go:embed tpl_question_readme.tpl
var TplQuestionReadme string

//go:embed tpl_question_go_test.tpl
var TplQuestionGoTest string

//go:embed tpl_question_js_test.tpl
var TplQuestionJsTest string

//go:embed tpl_record.tpl
var TplRecord string

//go:embed tpl_tag.tpl
var TplTag string

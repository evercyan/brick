package config

import (
	_ "embed" // ...
)

// TplQuestionReadme ...
//
//go:embed tpl_question_readme.tpl
var TplQuestionReadme string

// TplQuestionGoTest ...
//
//go:embed tpl_question_go_test.tpl
var TplQuestionGoTest string

// TplQuestionJsTest ...
//
//go:embed tpl_question_js_test.tpl
var TplQuestionJsTest string

// TplRecord ...
//
//go:embed tpl_record.tpl
var TplRecord string

// TplTag ...
//
//go:embed tpl_tag.tpl
var TplTag string

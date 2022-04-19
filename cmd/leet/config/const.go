package config

// 应用
const (
	App          = "leet"
	ConfigFile   = "leet.yaml"
	QuestionFile = "question.json"
)

// 符号
const (
	SymbolError   = "💣"
	SymbolNotice  = "🐶"
	SymbolSuccess = "🎉"
	SymbolTime    = "🕙"
)

// 答题语言配置
var (
	LangExt = map[string]string{
		"golang":     "go",
		"php":        "php",
		"python":     "py",
		"python3":    "py",
		"javascript": "js",
		"mysql":      "sql",
		"bash":       "sh",
	}
	LangList = []string{
		"golang", "php", "python", "python3", "javascript", "mysql", "bash",
	}
	LangMap = map[string]map[string]string{
		"golang": {
			"file":        "solution.go",
			"fileTpl":     "package solution\n\n%s",
			"testfile":    "solution_test.go",
			"testfileTpl": TplQuestionGoTest,
		},
		"javascript": {
			"file":        "solution.js",
			"fileTpl":     "%s\n\nmodule.exports = FuncToReplace;",
			"testfile":    "solution.test.js",
			"testfileTpl": TplQuestionJsTest,
		},
	}
	LangFile    = "solution.%s"
	LangFileTpl = "%s"
)

// LeetCode 接口地址
const (
	LeetCodeAllURL      = "https://leetcode-cn.com/api/problems/all/"  // 题目列表地址
	LeetCodeTagURL      = "https://leetcode-cn.com/tag/%s/problemset/" // 标签页面地址
	LeetCodeGraphqlURL  = "https://leetcode-cn.com/graphql"            // 题目数据地址
	LeetCodeQuestionURL = "https://leetcode-cn.com/problems/%s/"       // 题目页面地址
)

// 文件地址配置
const (
	QuestionPath = "question"
	CategoryPath = "category"
)

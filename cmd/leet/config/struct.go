package config

// ----------------------------------------------------------------

// Question ...
type Question struct {
	Fid    string          `json:"fid"`    // 显示 id
	Qid    int64           `json:"qid"`    // 实际 id
	Title  string          `json:"title"`  // 名称
	Slug   string          `json:"slug"`   // 标识
	Link   string          `json:"link"`   // 链接
	Level  QuestionLevel   `json:"level"`  // 困难度
	Detail *QuestionDetail `json:"detail"` // 详情
}

// ----------------------------------------------------------------

// QuestionDetail ...
type QuestionDetail struct {
	Title    string              `json:"title"`
	Content  string              `json:"content"`
	TagList  []map[string]string `json:"tag_list"`
	LangList []string            `json:"lang_list"`
	CodeMap  map[string]string   `json:"code_map"`
}

// ----------------------------------------------------------------

// QuestionLevel 题目难度
type QuestionLevel int

const (
	QuestionLevelEasy QuestionLevel = iota + 1
	QuestionLevelMiddle
	QuestionLevelHard
)

func (t QuestionLevel) String() string {
	switch t {
	case QuestionLevelEasy:
		return "简单"
	case QuestionLevelMiddle:
		return "中等"
	case QuestionLevelHard:
		return "困难"
	default:
		return ""
	}
}

package xtime

// 格式化占位符
const (
	FormatDate     = "ymd"
	FormatDateBar  = "y-m-d"
	FormatTime     = "y-m-d h:i:s"
	FormatTimeJoin = "ymdhis"
)

// ...
var formatMap = map[string]string{
	"y": "2006",
	"m": "01",
	"d": "02",
	"h": "15",
	"i": "04",
	"s": "05",
}

package internal

// Symbol 符号
type Symbol int

// 符号
const (
	SymbolNone Symbol = iota
	SymbolBegin
	SymbolSuccess
	SymbolFail
)

func (t Symbol) String() string {
	switch t {
	case SymbolBegin:
		return "🚀"
	case SymbolSuccess:
		return "🎉"
	case SymbolFail:
		return "🧨"
	default:
		return ""
	}
}

// ----------------------------------------------------------------

const (
	cmdSuffixA    = `awk '{a[NR]=$0}NR>1{print a[NR-1]}END{print substr(a[NR],0,length(a[NR])-1)}'`
	cmdSuffixB    = `sort | uniq -c | awk '{print "{\"k\": \""$2"\", \"v\": \""$1"\"},"}' | ` + cmdSuffixA
	cmdGetProject = `git remote -v | awk '{print $2}' | awk -F '/' '{print $NF}' | head -n 1`
	cmdGetLog     = `git log remotes/origin/master --pretty=format:"%cd|%h|%cn|%s" --date=format:"%Y-%m-%d %H:%M:%S"`
	cmdGetDate    = `cat %s | awk -F '|' '{print $1}' | awk '{print $1}' | ` + cmdSuffixB
	cmdGetWeek    = `cat %s | awk -F '|' '{print $1}' | awk '{system("date -j -f %%Y-%%m-%%d "$1" +%%w")}' | sort | uniq -c | awk '{print "{\"k\": \""$2"\", \"v\": \""$1"\"},"}' | sed 's/"k": "0"/"k": "Sun"/' | sed 's/"k": "1"/"k": "Mon"/' | sed 's/"k": "2"/"k": "Tue"/' | sed 's/"k": "3"/"k": "Wed"/' | sed 's/"k": "4"/"k": "Thu"/' | sed 's/"k": "5"/"k": "Fri"/' | sed 's/"k": "6"/"k": "Sat"/' | ` + cmdSuffixA
	cmdGetHour    = `cat %s | awk -F '|' '{print $1}' | awk '{print $2}' | awk -F ':' '{print $1}' | ` + cmdSuffixB
	cmdGetMin     = `cat %s | awk -F '|' '{print $1}' | awk '{print $2}' | awk -F ':' '{print $1":"$2}' | ` + cmdSuffixB
	cmdGetUser    = `cat %s | awk -F '|' '{print $3}' | ` + cmdSuffixB
	cmdGetWord    = `cat %s | sed 's/\"//g' | awk -F '|' '{print "\""$4"\","}' |` + cmdSuffixA
)

// analysisItems ...
var analysisItems = []map[string]string{
	{
		"title": "By Date",
		"type":  "bar",
		"key":   "date",
		"cmd":   cmdGetDate,
		"x":     "date",
		"y":     "num",
	},
	{
		"title": "By Hour",
		"type":  "bar",
		"key":   "hour",
		"cmd":   cmdGetHour,
		"x":     "hour",
		"y":     "num",
	},
	{
		"title": "By Week",
		"type":  "bar",
		"key":   "week",
		"cmd":   cmdGetWeek,
		"x":     "week",
		"y":     "num",
	},
	// {
	// 	"title": "By Min",
	// 	"type":  "line",
	// 	"key":   "min",
	// 	"cmd":   cmdGetMin,
	// 	"x":     "min",
	// 	"y":     "num",
	// },
	{
		"title": "By Committer",
		"type":  "pie",
		"key":   "user",
		"cmd":   cmdGetUser,
		"x":     "committer",
		"y":     "num",
	},
	{
		"title": "By Keyword",
		"type":  "wordcloud",
		"key":   "word",
		"cmd":   cmdGetWord,
		"x":     "keyword",
		"y":     "num",
	},
}

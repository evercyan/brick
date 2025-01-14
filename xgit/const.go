package xgit

// Client https://docs.github.com/cn/rest/reference/repos#contents
const (
	ApiURL         = "https://api.github.com/repos/%s/%s/contents/%s"
	TagURL         = "https://api.github.com/repos/%s/%s/tags"
	CdnURL         = "https://cdn.jsdelivr.net/gh/%s/%s/%s"
	DefaultMessage = "operate by brick"
)

package xregex

// MatchChinese ...
func MatchChinese(s string) []string {
	return match(patternChinese, s)
}

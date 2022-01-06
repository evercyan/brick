package xregexp

// MatchChinese ...
func MatchChinese(str string) []string {
	return match(patternChinese, str)
}

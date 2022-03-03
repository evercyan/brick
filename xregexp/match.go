package xregexp

// MatchChinese ...
func MatchChinese(s string) []string {
	return match(patternChinese, s)
}

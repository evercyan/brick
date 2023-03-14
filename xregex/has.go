package xregex

// HasIP ...
func HasIP(s string) bool {
	return HasIPV4(s) || HasIPV6(s)
}

// HasIPV4 ...
func HasIPV4(s string) bool {
	return has(patternIPV4, s)
}

// HasIPV6 ...
func HasIPV6(s string) bool {
	return has(patternIPV6, s)
}

// HasPhone ...
func HasPhone(s string) bool {
	return has(patternPhone, s)
}

// HasEmail ...
func HasEmail(s string) bool {
	return has(patternEmail, s)
}

// HasLink ...
func HasLink(s string) bool {
	return has(patternLink, s)
}

// HasDate ...
func HasDate(s string) bool {
	return has(patternDate, s)
}

// HasTime ...
func HasTime(s string) bool {
	return has(patternTime, s)
}

// HasChinese ...
func HasChinese(s string) bool {
	return has(patternChinese, s)
}

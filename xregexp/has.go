package xregexp

// HasIP ...
func HasIP(str string) bool {
	return HasIPV4(str) || HasIPV6(str)
}

// HasIPV4 ...
func HasIPV4(str string) bool {
	return has(patternIPV4, str)
}

// HasIPV6 ...
func HasIPV6(str string) bool {
	return has(patternIPV6, str)
}

// HasPhone ...
func HasPhone(str string) bool {
	return has(patternPhone, str)
}

// HasEmail ...
func HasEmail(str string) bool {
	return has(patternEmail, str)
}

// HasLink ...
func HasLink(str string) bool {
	return has(patternLink, str)
}

// HasDate ...
func HasDate(str string) bool {
	return has(patternDate, str)
}

// HasTime ...
func HasTime(str string) bool {
	return has(patternTime, str)
}

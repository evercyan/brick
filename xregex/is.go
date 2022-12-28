package xregex

// IsIPV4 ...
func IsIPV4(s string) bool {
	return is(patternIPV4, s)
}

// IsIPV6 ...
func IsIPV6(s string) bool {
	return is(patternIPV6, s)
}

// IsIP ...
func IsIP(s string) bool {
	return IsIPV4(s) || IsIPV4(s)
}

// IsMacAddress ...
func IsMacAddress(s string) bool {
	return is(patternMacAddress, s)
}

// IsPhone ...
func IsPhone(s string) bool {
	return is(patternPhone, s)
}

// IsEmail ...
func IsEmail(s string) bool {
	return is(patternEmail, s)
}

// IsLink ...
func IsLink(s string) bool {
	return is(patternLink, s)
}

// IsDate ...
func IsDate(s string) bool {
	return is(patternDate, s)
}

// IsTime ...
func IsTime(s string) bool {
	return is(patternTime, s)
}

// IsHexColor ...
func IsHexColor(s string) bool {
	return is(patternHexColor, s)
}

// IsIdcard ...
func IsIdcard(s string) bool {
	return is(patternIdcard, s)
}

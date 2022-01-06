package xregexp

// IsIPV4 ...
func IsIPV4(str string) bool {
	return is(patternIPV4, str)
}

// IsIPV6 ...
func IsIPV6(str string) bool {
	return is(patternIPV6, str)
}

// IsIP ...
func IsIP(str string) bool {
	return IsIPV4(str) || IsIPV4(str)
}

// IsMacAddress ...
func IsMacAddress(str string) bool {
	return is(patternMacAddress, str)
}

// IsPhone ...
func IsPhone(str string) bool {
	return is(patternPhone, str)
}

// IsEmail ...
func IsEmail(str string) bool {
	return is(patternEmail, str)
}

// IsLink ...
func IsLink(str string) bool {
	return is(patternLink, str)
}

// IsDate ...
func IsDate(str string) bool {
	return is(patternDate, str)
}

// IsTime ...
func IsTime(str string) bool {
	return is(patternTime, str)
}

// IsHexColor ...
func IsHexColor(str string) bool {
	return is(patternHexColor, str)
}

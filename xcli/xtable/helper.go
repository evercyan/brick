package xtable

// repeat ...
func repeat(char rune, num int) string {
	var s = make([]rune, num)
	for i := range s {
		s[i] = char
	}
	return string(s)
}

// getLentgh 获取实际显示长度
func getLentgh(r []rune) int {
	length := len(r)
	for _, v := range r {
		if v >= chineseCharset[0] && v <= chineseCharset[1] {
			length++
		}
	}
	return length
}

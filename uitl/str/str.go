package str

/**
 *	截取字符串
 */
func Substr(s string, pos, length int) string {
	if pos < 0 {
		return ""
	}
	runes := []rune(s)
	l := pos + length
	if pos > len(runes) {
		return ""
	}
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

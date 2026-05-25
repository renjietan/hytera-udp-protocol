package tools

// GetStringSize 获取字符串字的个数
func GetStringSize(str string) int {
	return len([]rune(str))
}

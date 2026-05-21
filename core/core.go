package core

func GetStringSize(str string) int {
	return len([]rune(str))
}

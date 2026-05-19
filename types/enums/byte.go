package enums

type Direction int

const (
	Backward Direction = iota // 向后：从 start 到末尾
	Forward                   // 向前：从 0 到 start（不包括 start）
)

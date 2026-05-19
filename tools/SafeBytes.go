package tools

import "github.com/renjietan/hytera-udp-protocol/types/enums"

type SafeBytes struct {
	data []byte
}

func NewSafeBytes(data []byte) *SafeBytes {
	return &SafeBytes{data: data}
}

// Slice 安全截取，返回新 SafeBytes 对象
func (sb *SafeBytes) Slice(start int, dir enums.Direction) *SafeBytes {
	if start < 0 {
		return &SafeBytes{data: []byte{}}
	}
	if start >= len(sb.data) {
		return &SafeBytes{
			data: sb.data,
		}
	}
	if dir == enums.Backward {
		return &SafeBytes{data: sb.data[start:]}
	} else {
		return &SafeBytes{data: sb.data[:start]}
	}
}

// GetByte 安全获取字节
func (sb *SafeBytes) GetByte(index int) byte {
	if index < 0 || index >= len(sb.data) {
		return 0
	}
	return sb.data[index]
}

// Bytes 返回底层切片（如果需要）
func (sb *SafeBytes) Bytes() []byte {
	return sb.data
}

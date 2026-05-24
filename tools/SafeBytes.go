package tools

import "github.com/renjietan/hytera-udp-protocol/types/enums"

type SafeBytes struct {
	Data []byte
}

func NewSafeBytes(data []byte) *SafeBytes {
	return &SafeBytes{Data: data}
}

// Slice 安全截取，返回新 SafeBytes 对象
func (sb *SafeBytes) Slice(start int, dir enums.Direction) *SafeBytes {
	if start < 0 {
		return &SafeBytes{Data: []byte{}}
	}
	if start >= len(sb.Data) {
		return &SafeBytes{
			Data: sb.Data,
		}
	}
	if dir == enums.Backward {
		return &SafeBytes{Data: sb.Data[start:]}
	} else {
		return &SafeBytes{Data: sb.Data[:start]}
	}
}

// GetByte 安全获取字节
func (sb *SafeBytes) GetByte(index int) byte {
	if index < 0 || index >= len(sb.Data) {
		return 0
	}
	return sb.Data[index]
}

// Bytes 返回底层切片（如果需要）
func (sb *SafeBytes) Bytes() []byte {
	return sb.Data
}

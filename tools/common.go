package tools

import (
	"github.com/renjietan/hytera-udp-protocol/types"
)

func ChunkByInterface[T types.ChunkInterface](data T, size int) []T {
	length := len(data)
	var res = []T{}
	for i := 0; i < length; i += size {
		end := i + size
		if end > length {
			end = length
		}
		res = append(res, data[i:end])
	}
	return res
}

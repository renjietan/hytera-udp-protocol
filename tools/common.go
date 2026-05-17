package tools

import (
	"encoding/binary"
	"errors"
)

func Uint2Bytes(val uint64, size int, IsBE bool) ([]byte, error) {
	buf := make([]byte, size)
	if size == 1 {
		buf[0] = byte(val)
		return buf, nil
	}
	if IsBE {
		switch size {
		case 2:
			binary.BigEndian.PutUint16(buf, uint16(val))
		case 4:
			binary.BigEndian.PutUint32(buf, uint32(val))
		case 8:
			binary.BigEndian.PutUint64(buf, val)
		default:
			return nil, errors.New("Uint2Bytes: 无效的字节尺寸")
		}
	} else {
		switch size {
		case 2:
			binary.LittleEndian.PutUint16(buf, uint16(val))
		case 4:
			binary.LittleEndian.PutUint32(buf, uint32(val))
		case 8:
			binary.LittleEndian.PutUint64(buf, val)
		default:
			return nil, errors.New("Uint2Bytes: 无效的字节尺寸")
		}
	}
	return buf, nil
}

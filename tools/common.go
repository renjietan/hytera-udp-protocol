package tools

import (
	"encoding/binary"
	"errors"
	"reflect"

	"github.com/renjietan/hytera-udp-protocol/types"
)

// GetRecursiveField 返回：针对需要统计长度的字段 进行计数，例如 [{ Name: "Length", Value: 10, Size: 2  }]
// params{t} 用于记录数据中，哪些数据是需要统计长度的，统计出长度后，全部放入 t 中
// params{params} 全量数据
func GetRecursiveField(params types.UdpRequest, t types.UdpRequest) types.UdpRequest {
	for _, v := range params {
		size := v.Size
		name := v.Name
		switch value := v.Value.(type) {
		case types.UdpRequest:
			//fmt.Println("递归：", name)
			t = GetRecursiveField(value, t)
		case int, []byte:
			if len(t) > 0 {
				lastIndex := len(t) - 1
				t[lastIndex].Value = t[lastIndex].Value.(int) + size
				for i := 0; i < lastIndex; i++ {
					t[i].Value = t[i].Value.(int) + size
				}
				//fmt.Println("计算1", t, name)
			}
		case nil:
			t = append(t, types.Item{
				Name:  name,
				Size:  size,
				Value: 0x00,
			})
			if len(t) > 1 {
				lastIndex := len(t) - 1
				for i := 0; i < lastIndex; i++ {
					t[i].Value = t[i].Value.(int) + size
				}
			}
			//fmt.Println("新增", t, name)
		}
	}
	return t
}

// Struct2Bytes  针对需要计算长度的字段，返回最终结果:
func Struct2Bytes(params types.UdpRequest, t []types.Item, res []byte) (types.UdpRequest, []byte) {
	for index, v := range params {
		switch v.Value.(type) {
		case types.UdpRequest:
			t, res = Struct2Bytes(params[index].Value.(types.UdpRequest), t, res)
		default:
			for i := 0; i < len(t); i++ {
				if t[i].Name == v.Name {
					if i == len(t)-1 {
						params[index].Value = t[i].Value.(int)
					} else {
						params[index].Value = t[i].Value.(int) + t[i].Size
					}
					t = t[1:]
				}
			}
			vType := reflect.TypeOf(params[index].Value).String()
			var vBytes []byte
			if vType == "[]uint8" {
				vBytes = params[index].Value.([]byte)
			} else if vType == "int" {
				vBytes, _ = Uint2Bytes(uint64(params[index].Value.(int)), params[index].Size, params[index].IsLe)
			}
			res = append(res, vBytes...)
		}
	}
	return t, res
}

func Uint2Bytes(val uint64, size int, IsLe bool) ([]byte, error) {
	buf := make([]byte, size)
	if size == 1 {
		buf[0] = byte(val)
		return buf, nil
	}
	if IsLe {
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
	} else {
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
	}
	return buf, nil
}

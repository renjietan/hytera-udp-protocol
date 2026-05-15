package core

import (
	"fmt"

	"github.com/renjietan/hytera-udp-protocol/options"
)

//var Login = options.UdpRequest{
//	"SrcID":   0xEE,
//	"DstID":   0xEE,
//	"Length":  nil,
//	"CRC":     0x00,
//	"Version": 0x00,
//	"UserID":  0x00,
//	"SAP":     0x01,
//	"Payload": options.UdpRequest{
//		"OptCode": 0x01,
//		"OptData": options.UdpRequest{
//			"Size":     nil,
//			"UserName": 5,
//		},
//	},
//}

var Login = options.UdpRequest{{
	Name:  "SrcID",
	Value: 0xEE,
	Size:  1,
}, {
	Name:  "DstID",
	Value: 0xEE,
	Size:  1,
}, {
	Name:  "Length",
	Value: nil,
	Size:  2,
}, {
	Name:  "CRC",
	Value: 0x00,
	Size:  2,
}, {
	Name:  "Version",
	Value: 0x00,
	Size:  1,
}, {
	Name:  "UserID",
	Value: 0x00,
	Size:  1,
}, {
	Name:  "SAP",
	Value: 0x01,
	Size:  1,
}, {
	Name: "Payload",
	Value: options.UdpRequest{{
		Name:  "OptCode",
		Value: 0x01,
		Size:  1,
	}, {
		Name: "OptData",
		Value: options.UdpRequest{{
			Name:  "Size",
			Value: nil,
			Size:  1,
		}, {
			Name:  "UserName",
			Value: 0xFF,
			Size:  10,
		}, {
			Name: "Password",
			Value: options.UdpRequest{{
				Name:  "Password1",
				Value: 0x00,
				Size:  4,
			}, {
				Name:  "Password2",
				Value: 0x00,
				Size:  5,
			}, {
				Name: "Password3",
				Value: options.UdpRequest{{
					Name:  "Password3-1",
					Value: nil,
					Size:  1,
				}, {
					Name:  "Password3-2",
					Value: 0xFF,
					Size:  5,
				}},
				Size: 0,
			}},
			Size: 0,
		}},
		Size: 0,
	}},
	Size: 0,
}}

// GetRecursiveField 返回：针对需要统计长度的字段 进行计数，例如 [{ Name: "Lenght", Value: 10, Size: 2  }]
func GetRecursiveField(params options.UdpRequest, t []options.Item) []options.Item {
	for _, v := range params {
		size := v.Size
		name := v.Name
		switch value := v.Value.(type) {
		case options.UdpRequest:
			fmt.Println("递归：", name)
			t = GetRecursiveField(value, t)
		case int:
			if len(t) > 0 {
				lastIndex := len(t) - 1
				t[lastIndex].Value = t[lastIndex].Value.(int) + size
				for i := 0; i < lastIndex; i++ {
					t[i].Value = t[i].Value.(int) + size
				}
				fmt.Println("计算", t, name)
			}
		case nil:
			t = append(t, options.Item{
				Name:  name,
				Size:  size,
				Value: 0x00,
			})
			fmt.Println("新增", t, name)
		}
	}
	return t
}

func SetRecursiveValue(params options.UdpRequest, t []options.Item) []options.Item {
	for index, v := range params {
		name := v.Name
		switch value := v.Value.(type) {
		case options.UdpRequest:
			t = SetRecursiveValue(value, t)
		default:
			for i := 0; i < len(t); i++ {
				if t[i].Name == name {
					if i == len(t)-1 {
						params[index].Value = t[i].Value.(int)
					} else {
						params[index].Value = t[i].Value.(int) + t[i].Size
					}
					t = t[1:]
				}
			}
		}
	}
	return t
}

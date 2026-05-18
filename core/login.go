package core

import (
	"github.com/renjietan/hytera-udp-protocol/options"
	"github.com/renjietan/hytera-udp-protocol/tools"
)

var ResultLogin = func(username string, userId int, alive bool) []byte {
	loginByte, err := TempLogin("admin", 11, true)
	if err != nil {
		return nil
	}
	return Struct2Bytes(loginByte)
}

// TempLogin Mortal 2026/5/18 16:04 初始化 login 所需字节，返回结构体
// username{string}: 用户名称
// userId{int}: 用户id
// alive{bool} 需要保活吗
// todo: 可能此处 userId 需要保存下来
var TempLogin = func(username string, userId int, alive bool) (options.UdpRequest, error) {
	res, err := TempBase(userId)
	if err != nil {
		return nil, err
	}
	bUsername := []byte(username)
	optCode := tools.Tern(alive == true, 0x05, 0x01)
	if optCode == 0x01 {
		res = append(res, options.Item{
			Name: "Payload",
			Value: options.UdpRequest{{
				Name:  "OptCode",
				Value: optCode,
				Size:  1,
			}, {
				Name: "OptData",
				Value: options.UdpRequest{{
					Name:  "Size",
					Value: len(username),
					Size:  1,
				}, {
					Name:  "UserName",
					Value: bUsername,
					Size:  len(bUsername),
				}},
				Size: 0,
			}},
			Size: 0,
		})
	} else {
		res = append(res, options.Item{
			Name: "Payload",
			Value: options.UdpRequest{{
				Name:  "OptCode",
				Value: optCode,
				Size:  1,
			}, {
				Name: "OptData",
				Value: options.UdpRequest{{
					Name:  "SuperviseInterval",
					Value: 3000,
					Size:  4,
				}, {
					Name:  "superviseCnt",
					Value: 3,
					Size:  2,
				}, {
					Name:  "Size",
					Value: len(username),
					Size:  1,
				}, {
					Name:  "UserName",
					Value: bUsername,
					Size:  len(bUsername),
				}},
				Size: 0,
			}},
			Size: 0,
		})
	}
	return res, nil
}

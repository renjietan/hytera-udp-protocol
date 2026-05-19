package auth

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// Login 登录;
// 说明:
//   - duration传入0, 默认心跳间隔为3s
//   - duration传入0，默认延时器为3s
var Login = func(username string, userId int, duration int) ([]byte, error) {
	tempByte, err := TempLogin(username, userId, duration)
	if err != nil {
		return nil, errors.New("Failed to insert into the login template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.Item{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

// TempLogin Mortal 2026/5/18 16:04 初始化 login 所需字节，返回结构体
var TempLogin = func(username string, userId int, duration int) (types.UdpRequest, error) {
	res, err := core.TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	bUsername, _ := tools.EncodeString(username, "UTF-16BE")
	optCode := tools.Tern(duration > 0, 0x05, 0x01)
	if optCode == 0x01 {
		res = append(res, types.Item{
			Name: "Payload",
			Value: types.UdpRequest{{
				Name:  "OptCode",
				Value: optCode,
				Size:  1,
			}, {
				Name: "OptData",
				Value: types.UdpRequest{{
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
		res = append(res, types.Item{
			Name: "Payload",
			Value: types.UdpRequest{{
				Name:  "OptCode",
				Value: optCode,
				Size:  1,
			}, {
				Name: "OptData",
				Value: types.UdpRequest{{
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

package auth

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// LoginInfoReq 对于双通道电台，用户可以获取登录成功后的具体信息
// 说明:
// - 包括但不限于: 通道所属设备适配器名称，心跳检测间隔，登录时长，最近一次交互时间等
// - 用户可以通过设备适配器名称来感知所属通道
var LoginInfoReq = func(username string, userId int) ([]byte, error) {
	tempByte, err := TLoginInfoReq(username, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TLoginInfoReq template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.Item{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TLoginInfoReq = func(username string, userId int) (types.UdpRequest, error) {
	res, err := core.TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	bUsername, _ := tools.EncodeString(username, "UTF-16BE")
	res = append(res, types.Item{
		Name: "Payload",
		Value: types.UdpRequest{{
			Name:  "OptCode",
			Value: 0x09,
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
	return res, nil
}

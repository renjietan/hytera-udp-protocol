package core_request_auth

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/core/request/types"
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
	res := request.Struct2BytesCode(tempByte)
	return res, nil
}

var TLoginInfoReq = func(username string, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	bUsername, _ := request.EncodeString(username, "UTF-16BE")
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode", // 操作码
			Value: 0x09,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "Size", // 用户名的字的个数，最大40个字
				Value: len(username),
				Size:  1,
			}, {
				Name:  "UserName", // UTF-16BE编码。最大长度40个字
				Value: bUsername,
				Size:  len(bUsername),
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

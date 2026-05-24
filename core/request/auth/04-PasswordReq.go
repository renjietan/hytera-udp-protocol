package core_request_auth

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// PasswordReq 某些业务需要验证密码之后才能够使用，用户用此消息发送密码请求
var PasswordReq = func(password string, userId int) ([]byte, error) {
	tempByte, err := TPassword(password, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TPassword template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.UdpRequestByteCodeItem{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TPassword = func(password string, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	bPwd, _ := tools.EncodeString(password, "UTF-16BE")
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x04,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "Size",
				Value: len(password),
				Size:  1,
			}, {
				Name:  "UserName",
				Value: bPwd,
				Size:  len(bPwd),
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

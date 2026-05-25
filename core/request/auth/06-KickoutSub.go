package core_request_auth

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// KickOutSub 踢出用户的订阅
var KickOutSub = func(username string, userId int) ([]byte, error) {
	tempByte, err := TKickOutSub(username, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TKickOutSub template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.UdpRequestByteCodeItem{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TKickOutSub = func(username string, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	bPwd, _ := tools.EncodeString(username, "UTF-16BE")
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode", // 操作码
			Value: 0x06,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "Size", // 用户名的字的个数，最大40个字
				Value: len(username),
				Size:  1,
			}, {
				Name:  "UserName", // UTF-16BE编码。最大长度40个字
				Value: bPwd,
				Size:  len(bPwd),
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

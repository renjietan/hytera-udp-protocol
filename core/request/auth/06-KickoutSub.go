package core_request_auth

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// KickOutSub 踢出用户的订阅
var KickOutSub = func(username string, userId int) ([]byte, error) {
	tempByte, err := TKickOutSub(username, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TKickOutSub template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.Item{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TKickOutSub = func(username string, userId int) (types.UdpRequest, error) {
	res, err := core.TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	bPwd, _ := tools.EncodeString(username, "UTF-16BE")
	res = append(res, types.Item{
		Name: "Payload",
		Value: types.UdpRequest{{
			Name:  "OptCode",
			Value: 0x06,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequest{{
				Name:  "Size",
				Value: len(username),
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

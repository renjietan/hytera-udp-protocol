package auth

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// Logout 用户主动断开与电台的连接时向电台发送登出请求
var Logout = func(username string, userId int) ([]byte, error) {
	tempByte, err := TempLogout(username, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the logout template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.Item{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TempLogout = func(username string, userId int) (types.UdpRequest, error) {
	res, err := core.TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	bUsername, _ := tools.EncodeString(username, "UTF-16BE")
	res = append(res, types.Item{
		Name: "Payload",
		Value: types.UdpRequest{{
			Name:  "OptCode",
			Value: 0x02,
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

package core_request_auth

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// LogoutReq Logout 主动断开与电台的连接时向电台发送登出请求
var LogoutReq = func(username string, userId int) ([]byte, error) {
	tempByte, err := TLogout(username, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TLogout template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.UdpRequestByteCodeItem{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TLogout = func(username string, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	bUsername, _ := tools.EncodeString(username, "UTF-16BE")
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x02,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
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

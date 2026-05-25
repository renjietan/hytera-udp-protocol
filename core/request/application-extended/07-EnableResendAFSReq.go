package core_request_application_extended

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/core/request/types"
)

// EnableResendAFSReq 使能应用功能业务确认重发机制
var EnableResendAFSReq = func(Version, userId int) ([]byte, error) {
	tempByte, err := TEnableResendAFSReq(Version, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TEnableResendAFSReq template: " + err.Error())
	}
	res := request.Struct2BytesCode(tempByte)
	return res, nil
}

var TEnableResendAFSReq = func(Version, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x03)
	if err != nil {
		return nil, err
	}
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x07,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "Version", // 确认重发机制版本号
				Value: Version,   // 例如: 若版本号为V0.1，则该字段表示为0x01
				Size:  2,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

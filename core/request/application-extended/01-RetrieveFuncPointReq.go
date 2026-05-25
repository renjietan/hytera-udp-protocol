package core_request_application_extended

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// RetrieveFuncPointReq 查找功能点
var RetrieveFuncPointReq = func(FuncPoint, userId int) ([]byte, error) {
	tempByte, err := TRetrieveFuncPointReq(FuncPoint, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TRetrieveFuncPointReq template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.UdpRequestByteCodeItem{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TRetrieveFuncPointReq = func(FuncPoint, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x03)
	if err != nil {
		return nil, err
	}
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x01,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "FuncPoint", // 功能点枚举，0为无效值
				Value: FuncPoint,
				Size:  2,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

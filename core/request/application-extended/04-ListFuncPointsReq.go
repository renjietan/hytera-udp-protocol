package core_request_application_extended

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// ListFuncPointsReq User向目标机发送功能点列表请求
var ListFuncPointsReq = func(ListType, userId int) ([]byte, error) {
	tempByte, err := TListFuncPointsReq(ListType, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TListFuncPointsReq template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.UdpRequestByteCodeItem{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TListFuncPointsReq = func(ListType, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x03)
	if err != nil {
		return nil, err
	}
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x04,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name: "ListType",
				// 0x01-所有正在运行的功能点
				// 0x02-所有禁用的功能点。
				Value: ListType,
				Size:  2,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

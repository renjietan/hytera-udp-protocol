package core_request_application_extended

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// RetrieveFuncPointReq 查找功能点
var RetrieveFuncPointReq = func(FuncPoint, userId int) ([]byte, error) {
	tempByte, err := TRetrieveFuncPointReq(FuncPoint, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TRetrieveFuncPointReq template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.Item{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TRetrieveFuncPointReq = func(FuncPoint, userId int) (types.UdpRequest, error) {
	res, err := core.TempBase(userId, 0x03)
	if err != nil {
		return nil, err
	}
	res = append(res, types.Item{
		Name: "Payload",
		Value: types.UdpRequest{{
			Name:  "OptCode",
			Value: 0x01,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequest{{
				Name:  "FuncPoint",
				Value: FuncPoint,
				Size:  2,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

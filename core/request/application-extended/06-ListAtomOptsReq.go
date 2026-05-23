package core_request_application_extended

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// ListAtomOptsReq 功能点支持的原子操作列表请求
var ListAtomOptsReq = func(FuncPoint, userId int) ([]byte, error) {
	tempByte, err := TListAtomOptsReq(FuncPoint, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TListAtomOptsReq template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.Item{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TListAtomOptsReq = func(FuncPoint, userId int) (types.UdpRequest, error) {
	res, err := core.TempBase(userId, 0x03)
	if err != nil {
		return nil, err
	}
	res = append(res, types.Item{
		Name: "Payload",
		Value: types.UdpRequest{{
			Name:  "OptCode",
			Value: 0x06,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequest{{
				Name:  "FuncPoint", // 功能点枚举，0为无效值
				Value: FuncPoint,   // 例如: 见 2.5.3功能点 0x0008-波形管理功能点(功能点枚举，0为无效值)
				Size:  2,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

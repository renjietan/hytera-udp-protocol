package core_request_application_extended

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// RunFuncPointReq 运行功能点请求
var RunFuncPointReq = func(FuncPoint, RunType, userId int) ([]byte, error) {
	tempByte, err := TRunFuncPointReq(FuncPoint, RunType, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TRunFuncPointReq template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.UdpRequestByteCodeItem{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TRunFuncPointReq = func(FuncPoint, RunType, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x03)
	if err != nil {
		return nil, err
	}
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x03,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "FuncPoint", // 功能点枚举，0为无效值
				Value: FuncPoint,   // 例如: 见 2.5.3功能点 0x0008-波形管理功能点(功能点枚举，0为无效值)
				Size:  2,
			}, {
				Name: "RunType",
				// 0x01-运行
				// 0x02-禁用
				Value: RunType,
				Size:  1,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

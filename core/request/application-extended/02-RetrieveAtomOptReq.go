package core_request_application_extended

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/core/request/types"
)

// RetrieveAtomOptReq 查找原子操作请求
var RetrieveAtomOptReq = func(FuncPoint, OptCode, OptType, userId int) ([]byte, error) {
	tempByte, err := TRetrieveAtomOptReq(FuncPoint, OptCode, OptType, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TRetrieveAtomOptReq template: " + err.Error())
	}
	res := request.Struct2BytesCode(tempByte)
	return res, nil
}

var TRetrieveAtomOptReq = func(FuncPoint, OptCode, OptType, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x03)
	if err != nil {
		return nil, err
	}
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x02,
			Size:  2,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "FuncPoint", // 功能点枚举，0为无效值
				Value: FuncPoint,   // 例如: 见 2.5.3功能点 0x0008-波形管理功能点(功能点枚举，0为无效值)
				Size:  2,
			}, {
				Name:  "OptType", // 操作类型
				Value: OptType,   //  例如: 见 2.5.2.1操作类型; 0x10-波形操作(waveform)
				Size:  1,
			}, {
				Name:  "OptCode", // 操作码。原子操作的唯一标识
				Value: OptCode,   // 例如: 见 2.5.2.10 (0x10)波形操作;
				Size:  2,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

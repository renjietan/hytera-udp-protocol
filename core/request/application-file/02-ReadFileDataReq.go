package core_request_application_file

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// ReadFileDataReq 读文件内容请求
var ReadFileDataReq = func(PacketNum, userId int) ([]byte, error) {
	tempByte, err := TReadFileDataReq(PacketNum, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TReadFileDataReq template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.UdpRequestByteCodeItem{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TReadFileDataReq = func(PacketNum, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x04)
	if err != nil {
		return nil, err
	}
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x02,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "PacketNum", // PacketNum 文件包编号
				Value: PacketNum,   // 例如: 0，1，2，3
				Size:  2,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

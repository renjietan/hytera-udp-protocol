package core_request_application_extended

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/core/request/types"
)

// ReadFuncPointsListReq 读取功能点列表请求
var ReadFuncPointsListReq = func(SegmentIndex, userId int) ([]byte, error) {
	tempByte, err := TReadFuncPointsListReq(SegmentIndex, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TReadFuncPointsListReq template: " + err.Error())
	}
	res := request.Struct2BytesCode(tempByte)
	return res, nil
}

var TReadFuncPointsListReq = func(SegmentIndex, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x03)
	if err != nil {
		return nil, err
	}
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x05,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "SegmentIndex", // 段编号，从1开始
				Value: SegmentIndex,
				Size:  1,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

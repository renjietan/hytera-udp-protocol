package core_request_application_file

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// ReadFileInfoReq 读文件头信息请求
var ReadFileInfoReq = func(FileName string, userId int) ([]byte, error) {
	tempByte, err := TReadFileInfoReq(FileName, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TReadFileInfoReq template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.UdpRequestByteCodeItem{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TReadFileInfoReq = func(FileName string, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x04)
	if err != nil {
		return nil, err
	}
	_FileName := []byte(FileName)
	FileNameLen := tools.GetStringSize(FileName)
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x01,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "NameSize",
				Value: FileNameLen,
				Size:  4,
			}, {
				Name:  "FileName",
				Value: _FileName,      // 文件名（含相对路径）
				Size:  len(_FileName), // 例如：./FPGA/fh.bin
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

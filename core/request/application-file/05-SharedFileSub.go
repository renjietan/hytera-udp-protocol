package core_request_application_file

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/core/request/types"
)

// SharedFileSub 共享文件订阅
var SharedFileSub = func(FileName string, userId int) ([]byte, error) {
	tempByte, err := TSharedFileSub(FileName, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TSharedFileSub template: " + err.Error())
	}
	res := request.Struct2BytesCode(tempByte)
	return res, nil
}

var TSharedFileSub = func(FileName string, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x04)
	if err != nil {
		return nil, err
	}
	FileNameLen := request.GetStringSize(FileName)
	FileNameByte := []byte(FileName)
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x05,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "NameSize",
				Value: FileNameLen,
				Size:  4,
			}, {
				Name:  "FileName",
				Value: FileNameByte,
				Size:  len(FileNameByte),
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

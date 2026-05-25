package core_request_application_file

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// WriteFileDataReq 写文件内容请求
var WriteFileDataReq = func(params types.UdpWriteFileDataReq, userId int) ([]byte, error) {
	var res = []byte{}
	PacketCrc := params.PacketCrc                          // 每包检验和
	PacketSize := params.PacketSize                        // 单个包的大小
	FileBody := params.FileBody                            // 文件内容（字节）
	Chunks := tools.ChunkByInterface(FileBody, PacketSize) // 分包

	for index, value := range Chunks {
		tempByte, err := TWriteFileDataReq(PacketCrc, index, PacketSize, value, userId)
		if err != nil {
			return nil, errors.New("Failed to insert into the TWriteFileDataReq template: " + err.Error())
		}
		recordField := tools.GetRecursiveField(tempByte, []types.UdpRequestByteCodeItem{})
		_, item := tools.Struct2Bytes(tempByte, recordField, []byte{})
		res = append(res, item...)
	}
	return res, nil
}

var TWriteFileDataReq = func(PacketCrc, PacketNum, PacketSize int, Data []byte, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x04)
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
				Name:  "PacketCrc", // 整个文件CRC校验位
				Value: PacketCrc,
				Size:  1,
			}, {
				Name:  "PacketNum", // 文件分发的个数
				Value: PacketNum,
				Size:  2,
			}, {
				Name:  "PacketSize", // 文件长度
				Value: PacketSize,
				Size:  4,
			}, {
				// TODO: 文件名称的字数 还是 文件名称的字节数
				Name:  "Data", // 文件名长度
				Value: Data,
				Size:  len(Data),
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

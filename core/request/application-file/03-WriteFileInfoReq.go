package core_request_application_file

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// WriteFileInfoReq 写文件头信息请求
var WriteFileInfoReq = func(params types.UdpWriteFileInfoReq, userId int) ([]byte, error) {

	tempByte, err := TWriteFileInfoReq(params, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TWriteFileInfoReq template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.UdpRequestByteCodeItem{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TWriteFileInfoReq = func(params types.UdpWriteFileInfoReq, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x04)
	if err != nil {
		return nil, err
	}
	FileBody := params.FileBody   // 文件内容（字节）
	ChunkSize := params.ChunkSize // 每包字节长度
	FileName := params.FileName   // 文件名（含相对路径）,例如：./FPGA/fh.bin
	FileNameLen := tools.GetStringSize(FileName)
	FileCRC := params.FileCRC
	Chunks := tools.ChunkByInterface(FileBody, ChunkSize) // 分包，返回数组
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode",
			Value: 0x03,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				Name:  "FileCRC", // 整个文件CRC校验位
				Value: FileCRC,
				Size:  1,
			}, {
				Name:  "PacketCnt", // 文件分发的个数
				Value: len(Chunks),
				Size:  2,
			}, {
				Name:  "FileSize", // 文件长度
				Value: len(FileBody),
				Size:  4,
			}, {
				// TODO: 文件名称的字数 还是 文件名称的字节数
				Name:  "NameLen", // 文件名长度
				Value: FileNameLen,
				Size:  4,
			}, {
				Name:  "FileName", // 文件名（含相对路径）
				Value: FileName,   // 例如：./FPGA/fh.bin
				Size:  len(FileName),
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

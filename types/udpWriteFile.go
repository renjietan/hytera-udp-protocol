package types

type ChunkInterface interface {
	~string | ~[]byte
}

type UdpWriteFileInfoReq struct {
	FileCRC   int    // 整个文件CRC校验位
	FileName  string // 文件名（含相对路径）,例如：./FPGA/fh.bin
	FileBody  []byte // 文件内容
	ChunkSize int    // 每包字节长度
}

type UdpWriteFileDataReq struct {
	PacketCrc  int    // 每包检验和
	PacketSize int    // 单个包大小（默认为500，最后一包小于500）
	FileBody   []byte // 文件内容
}

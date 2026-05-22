package types

type UdpRequest []Item

type Item struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
	Size  int    `json:"size"`
	IsLe  bool   `json:"isLE"`
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

type ApplicationInfo struct {
	// 功能点枚举，0为无效值
	// 例如: 见 2.5.3功能点 0x0008-波形管理功能点(功能点枚举，0为无效值)
	FuncPoint int
	// 所属通道号，0 代表通用功能点；从1开始，表示通道一、通道二等等,
	// 例如:0x00
	ChannelNo int
	// 事务ID。用于实现并发业务:
	// 例如: rand.Intn(30000) 0-30000随机数
	EventID int
	// 事务随机码。用于标识消息的唯一性
	// 例如: rand.Intn(30000) 0-30000随机数
	EventRN int
	AtomOpt struct {
		// 操作类型
		// 例如: 见 2.5.2.1操作类型; 0x10-波形操作(waveform)
		OptType int
		// 操作码。原子操作的唯一标识
		// 例如: 见 2.5.2.10 (0x10)波形操作;
		OptCode int
		// 调用类型（高四位表示方向，低四位表示类型）。
		// 请求(REQ): 0x01
		// 订阅(SUB): 0x02-SUB
		// 通知(NTY): 0x03-NTY
		// 应答(ACK): 0x04-ACK
		CallType int
		// 调用参数
		// 例如: 见(2.5.2.10.7) (0x0007) 子网参数配置
		CallPara UdpRequest
	}
}

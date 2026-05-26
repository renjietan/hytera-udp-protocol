package types

type UdpResponseBytesCode struct {
	SrcID   UdpResponseByteCodeItem
	DstID   UdpResponseByteCodeItem
	Length  UdpResponseByteCodeItem
	CRC     UdpResponseByteCodeItem
	Version UdpResponseByteCodeItem
	UserID  UdpResponseByteCodeItem
	SAP     UdpResponseByteCodeItem
	Payload any
}

type Payload struct {
	OptCode UdpResponseByteCodeItem // 1 Byte
	OptData any
}

type UdpResponseByteCodeItem struct {
	Value interface{}
	Size  int
	Bind  UdpResponseBindStruct
}

type AdapterInfoItem struct {
	// 设备适配器所属通道号, 0 代表不区分；从1开始，表示通道一、通道二等等。
	ChannelNo UdpResponseByteCodeItem // 1 Byte
	// 设备适配器类型：
	// 0x00    无效
	// 0x01    以太网
	// 0x02    串口
	// 0x03    被覆线
	AdapterType UdpResponseByteCodeItem // 1 Byte
	// 设备适配器IPv4地址。AdapterType为0x01时有效
	Ipv4Addr UdpResponseByteCodeItem // 4 Byte
	// 设备适配器IPv4掩码。AdapterType为0x01时有效
	Ipv4Mask UdpResponseByteCodeItem // 4 Byte
	// 设备适配器名称字的个数，最大255个字
	Size UdpResponseByteCodeItem // 1 Byte
	// 设备适配器名称: UTF-16BE编码
	AdapterName UdpResponseByteCodeItem // n Byte
}

type UdpResponseBindValueType string

const (
	UdpResponseBindStructInt    UdpResponseBindValueType = "int"
	UdpResponseBindStructString UdpResponseBindValueType = "string"
	UdpResponseBindStructRange  UdpResponseBindValueType = "Range"
)

type UdpResponseBindFileName string

const (
	UdpResponseBindStructSize  UdpResponseBindFileName = "Size"
	UdpResponseBindStructValue UdpResponseBindFileName = "Value"
)

type UdpResponseBindStruct struct {
	Path      string
	FieldName UdpResponseBindFileName
	//Callback  func(params ...any)
	Value     any
	ValueType UdpResponseBindValueType // 对应字段的 类型
	//Action    UdpResponseBindAction
}

//type UdpApplicationResponse struct {
//	SrcID   UdpBaseResponse
//	DstID   UdpBaseResponse
//	Length  UdpBaseResponse
//	CRC     UdpBaseResponse
//	Version UdpBaseResponse
//	UserID  UdpBaseResponse
//	SAP     UdpBaseResponse
//	OptCode UdpBaseResponse
//	OptData struct {
//		Value struct {
//			FuncPoint UdpBaseResponse
//			ChannelNo UdpBaseResponse
//			EventID   UdpBaseResponse
//			EventRN   UdpBaseResponse
//			AtomOpt   struct {
//				OptType  UdpBaseResponse
//				OptCode2 UdpBaseResponse
//				CallType UdpBaseResponse
//				ParaSize UdpBaseResponse
//				CallPara UdpBaseResponse
//			}
//		}
//	}
//}
//
//type UdpApplicationExtendedResponse struct {
//	SrcID   UdpBaseResponse
//	DstID   UdpBaseResponse
//	Length  UdpBaseResponse
//	CRC     UdpBaseResponse
//	Version UdpBaseResponse
//	UserID  UdpBaseResponse
//	SAP     UdpBaseResponse
//	OptCode UdpBaseResponse
//	OptData UdpBaseResponse
//}
//
//type UdpFileResponse struct {
//	SrcID   UdpBaseResponse
//	DstID   UdpBaseResponse
//	Length  UdpBaseResponse
//	CRC     UdpBaseResponse
//	Version UdpBaseResponse
//	UserID  UdpBaseResponse
//	SAP     UdpBaseResponse
//	OptCode UdpBaseResponse
//	OptData UdpBaseResponse
//}

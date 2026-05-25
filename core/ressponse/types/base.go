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
	Bind  UdpResonseBindStruct
}

type UdpResonseBindStruct struct {
	Path      string
	Callback  func(params ...any)
	Value     interface{}
	FieldName string
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

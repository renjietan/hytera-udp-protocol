package types

type UdpAuthResponse struct {
	SrcID   UdpBaseResponse
	DstID   UdpBaseResponse
	Length  UdpBaseResponse
	CRC     UdpBaseResponse
	Version UdpBaseResponse
	UserID  UdpBaseResponse
	SAP     UdpBaseResponse
	OptCode UdpBaseResponse
	OptData UdpBaseResponse
}

type UdpApplicationResponse struct {
	SrcID   UdpBaseResponse
	DstID   UdpBaseResponse
	Length  UdpBaseResponse
	CRC     UdpBaseResponse
	Version UdpBaseResponse
	UserID  UdpBaseResponse
	SAP     UdpBaseResponse
	OptCode UdpBaseResponse
	OptData struct {
		Value struct {
			FuncPoint UdpBaseResponse
			ChannelNo UdpBaseResponse
			EventID   UdpBaseResponse
			EventRN   UdpBaseResponse
			AtomOpt   struct {
				OptType  UdpBaseResponse
				OptCode2 UdpBaseResponse
				CallType UdpBaseResponse
				ParaSize UdpBaseResponse
				CallPara UdpBaseResponse
			}
		}
	}
}

type UdpApplicationExtendedResponse struct {
	SrcID   UdpBaseResponse
	DstID   UdpBaseResponse
	Length  UdpBaseResponse
	CRC     UdpBaseResponse
	Version UdpBaseResponse
	UserID  UdpBaseResponse
	SAP     UdpBaseResponse
	OptCode UdpBaseResponse
	OptData UdpBaseResponse
}

type UdpFileResponse struct {
	SrcID   UdpBaseResponse
	DstID   UdpBaseResponse
	Length  UdpBaseResponse
	CRC     UdpBaseResponse
	Version UdpBaseResponse
	UserID  UdpBaseResponse
	SAP     UdpBaseResponse
	OptCode UdpBaseResponse
	OptData UdpBaseResponse
}

type UdpBaseResponse struct {
	Value interface{}
	Size  int
}

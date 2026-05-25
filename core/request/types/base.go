package types

type UdpRequestBytesCode []UdpRequestByteCodeItem

type UdpRequestByteCodeItem struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
	Size  int    `json:"size"`
	IsLe  bool   `json:"isLE"`
}

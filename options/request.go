package options

type UdpRequest []Item

type Item struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
	Size  int    `json:"size"`
	IsBE  bool   `json:"isBE"`
}

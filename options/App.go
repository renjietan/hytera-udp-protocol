package options

import "net"

type Callback func(addr *net.UDPAddr, data string, err error)
type App struct {
	Host        string   `json:"host"`
	Port        string   `json:"port"`
	OnMsgFunc   Callback `json:"onMsgFunc"`
	OnCloseFunc Callback `json:"onCloseFunc"`
	OnErrorFunc Callback `json:"onErrorFunc"`
}

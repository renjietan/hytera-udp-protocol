package options

import "net"

type Callback func(addr *net.UDPAddr, data string, err error)
type App struct {
	Host        string   `yaml:"host"`
	Port        string   `json:"port"`
	RHost       string   `json:"rHost"`
	RPort       string   `json:"rPort"`
	OnMsgFunc   Callback `json:"onMsgFunc"`
	OnCloseFunc Callback `json:"onCloseFunc"`
	OnErrorFunc Callback `json:"onErrorFunc"`
}

package types

import "time"

type Callback func(data Envelope)
type App struct {
	Host        string        `yaml:"host"`
	Port        string        `json:"port"`
	Duration    time.Duration `yaml:"duration"`
	OnMsgFunc   Callback      `json:"onMsgFunc"`
	OnCloseFunc Callback      `json:"onCloseFunc"`
	OnErrorFunc Callback      `json:"onErrorFunc"`
}

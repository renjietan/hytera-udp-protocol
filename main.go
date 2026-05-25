package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/renjietan/hytera-udp-protocol/types"
)

func main() {
	client, err := NewUdpClient(&types.App{
		Host:     "127.0.0.1",
		Port:     "3333",
		Duration: time.Second * 3,
		OnMsgFunc: func(data types.Envelope) {
			fmt.Printf("msg event: %s; data: %#v; message: %s; code: %d; address: %v\n", data.Event, data.Data, data.Message, data.Code, data.Address)
		},
		OnErrorFunc: func(data types.Envelope) {
			fmt.Printf("error event: %s; data: %#v; message: %s; code: %d; address: %v\n", data.Event, data.Data, data.Message, data.Code, data.Address)
		},
		OnCloseFunc: func(data types.Envelope) {
			fmt.Printf("close event: %s; data: %#v; message: %s; code: %d; address: %v\n", data.Event, data.Data, data.Message, data.Code, data.Address)
		},
	})
	if err != nil {
		return
	}
	client.Login("8.135.10.183", 55022, "admin", 11, true)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 关闭应用程序
	fmt.Println("<-quit")
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := c.; err != nil {
	//
	//}
}

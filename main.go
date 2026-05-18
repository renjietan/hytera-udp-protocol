package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/renjietan/hytera-udp-protocol/options"
)

func main() {
	client, err := NewUdpClient(&options.App{
		Host:     "127.0.0.1",
		Port:     "3333",
		Duration: time.Second * 3,
		OnMsgFunc: func(data options.Envelope) {
			fmt.Println("onMsg：", data)
		},
		OnErrorFunc: func(data options.Envelope) {
			fmt.Println("onMsg：", data)
		},
		OnCloseFunc: func(data options.Envelope) {
			fmt.Println("onMsg：", data)
		},
	})
	client.Login("8.135.10.183", 57861, "admin", 11, true)
	if err != nil {
		return
	}

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

package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/renjietan/hytera-udp-protocol/options"
)

func main() {
	//loginByte, _ := core.TempLogin("admin", 11, true)
	//t := tools.GetRecursiveField(loginByte, []options.Item{})
	//tt, b := tools.Struct2Bytes(loginByte, t, []byte{})
	//marshal, err := json.Marshal(loginByte)
	//if err != nil {
	//	return
	//}
	//fmt.Println("结果", string(marshal))

	_, err := NewUdpClient(&options.App{
		Host:  "127.0.0.1",
		Port:  "3333",
		RHost: "8.135.10.183",
		RPort: "55022",
		OnMsgFunc: func(addr *net.UDPAddr, data any, err error) {
			fmt.Println("onMsg：", addr, data, err)
		},
		OnErrorFunc: func(addr *net.UDPAddr, data any, err error) {
			fmt.Println("onError：", addr, data, err)
		},
		OnCloseFunc: func(addr *net.UDPAddr, data any, err error) {
			fmt.Println("onClose：", addr, data, err)
		},
	})
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

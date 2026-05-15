package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/renjietan/hytera-udp-protocol/options"
)

func main() {
	var sw sync.WaitGroup
	sw.Add(1)
	c, err := NewUdpClient(&options.App{
		Host: "8.135.10.183",
		Port: "55022",
		OnMsgFunc: func(addr *net.UDPAddr, data string, err error) {
			fmt.Println(addr, data, err)
		},
		OnErrorFunc: func(addr *net.UDPAddr, data string, err error) {
			fmt.Println(addr, data, err)
		},
		OnCloseFunc: func(addr *net.UDPAddr, data string, err error) {
			fmt.Println(addr, data, err)
			sw.Done()
		},
	})
	fmt.Println("c=================", c)
	fmt.Println("err=================", err)
	sw.Wait()
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//// 关闭应用程序
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := c.; err != nil {
	//
	//}
}

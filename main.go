package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/options"
)

func main() {
	t := core.GetRecursiveField(core.Login, []options.Item{})
	fmt.Println("长度统计结果：", t)
	core.SetRecursiveValue(core.Login, t)
	marshal, err := json.Marshal(core.Login)
	if err != nil {
		return
	}
	fmt.Println("结果", string(marshal))
	//for _, item := range core.Login {
	//	fmt.Println("name:", item.Name, "size", item.Size)
	//}
	c, err := NewUdpClient(&options.App{
		Host:  "127.0.0.1",
		Port:  "3333",
		RHost: "8.135.10.183",
		RPort: "55022",
		OnMsgFunc: func(addr *net.UDPAddr, data string, err error) {
			fmt.Println(addr, data, err)
		},
		OnErrorFunc: func(addr *net.UDPAddr, data string, err error) {
			fmt.Println(addr, data, err)
		},
		OnCloseFunc: func(addr *net.UDPAddr, data string, err error) {
			fmt.Println(addr, data, err)
		},
	})
	//udpAddr, err := net.ResolveUDPAddr("udp", "8.135.10.183:55022")
	//if err != nil {
	//	return
	//}
	//err2 := c.Send("hello world")
	//if err2 != nil {
	//	return
	//}
	time.Sleep(1000 * time.Second)
	fmt.Println("c=================", c)
	fmt.Println("err=================", err)
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

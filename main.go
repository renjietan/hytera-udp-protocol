package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	client2 "github.com/renjietan/hytera-udp-protocol/client"
	"github.com/renjietan/hytera-udp-protocol/client/types"
	response "github.com/renjietan/hytera-udp-protocol/core/ressponse"
	core_response_auth "github.com/renjietan/hytera-udp-protocol/core/ressponse/auth"
	types2 "github.com/renjietan/hytera-udp-protocol/core/ressponse/types"
)

func main() {
	client, err := client2.NewUdpClient(&types.App{
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
	b := []byte{0xee, 0xee, 0x0, 0x17, 0x0, 0x0, 0x0, 0xb, 0x1, 0x5, 0x05,
		0x0, 0x61, 0x0, 0x64, 0x0, 0x6d, 0x0, 0x69, 0x0, 0x6e, 0x0c}
	lTemp := core_response_auth.KickOutUserNtyRes()
	response.ByteCode2Stuct(b, &lTemp, "", map[string]types2.UdpResonseBindStruct{})
	fmt.Println(lTemp)
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

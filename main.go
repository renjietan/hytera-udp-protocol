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
	//testA := []byte{0xee, 0xee, 0x0, 0x17, 0x0, 0x0, 0x0, 0xb, 0x1, 0x5, 0x01, 0x0b}
	//testInt := []byte{0xee, 0xee, 0x0, 0x17, 0x0, 0x0, 0x0, 0xb, 0x1, 0x5, 0x05,
	//	0x0, 0x61, 0x0, 0x64, 0x0, 0x6d, 0x0, 0x69, 0x0, 0x6e, 0x0c}
	testSlice := []byte{0xee, 0xee, 0x0, 0x17, 0x0, 0x0, 0x0, 0xb, 0x1, 0x5,
		0x00,       // Result
		0x00, 0x05, // ListCount

		0x00,                   // ChannelNo
		0x01,                   // AdapterType
		0x02, 0x03, 0x02, 0x03, // Ipv4Addr
		0x04, 0x05, 0x04, 0x05, // Ipv4Mask
		0x05,                                                  // Size
		0x0, 0x61, 0x0, 0x64, 0x0, 0x6d, 0x0, 0x69, 0x0, 0x6e, // AdapterName

		0x01,                   // ChannelNo
		0x02,                   // AdapterType
		0x03, 0x04, 0x03, 0x04, // Ipv4Addr
		0x05, 0x06, 0x05, 0x06, // Ipv4Mask
		0x05,                                                  // Size
		0x0, 0x61, 0x0, 0x64, 0x0, 0x6d, 0x0, 0x69, 0x0, 0x6e, // AdapterName

		0x02,                   // ChannelNo
		0x03,                   // AdapterType
		0x04, 0x05, 0x04, 0x05, // Ipv4Addr
		0x06, 0x07, 0x06, 0x07, // Ipv4Mask
		0x05,                                                  // Size
		0x0, 0x61, 0x0, 0x64, 0x0, 0x6d, 0x0, 0x69, 0x0, 0x6e, // AdapterName

		0x03,                   // ChannelNo
		0x04,                   // AdapterType
		0x05, 0x06, 0x05, 0x06, // Ipv4Addr
		0x07, 0x08, 0x07, 0x08, // Ipv4Mask
		0x05,                                                  // Size
		0x0, 0x61, 0x0, 0x64, 0x0, 0x6d, 0x0, 0x69, 0x0, 0x6e, // AdapterName

		0x04,                   // ChannelNo
		0x05,                   // AdapterType
		0x06, 0x07, 0x06, 0x07, // Ipv4Addr
		0x08, 0x09, 0x08, 0x09, // Ipv4Mask
		0x05,                                                  // Size
		0x0, 0x61, 0x0, 0x64, 0x0, 0x6d, 0x0, 0x69, 0x0, 0x6e, // AdapterName
	}

	//t := core_response_auth.LoginAckRes()
	//response.ByteCode2Stuct(testA, &t, "", map[string]types2.UdpResponseBindStruct{})
	//fmt.Printf("t: %#v\n", t)
	//fmt.Println("===================================================================")
	//
	//tInt := core_response_auth.KickOutUserNtyRes()
	//response.ByteCode2Stuct(testInt, &tInt, "", map[string]types2.UdpResponseBindStruct{})
	//fmt.Printf("testInt: %#v\n", tInt)
	//fmt.Println("===================================================================")

	tSlice := core_response_auth.NewAdapterInfoAckRes()
	response.ByteCode2Stuct(testSlice, &tSlice, "", map[string]types2.UdpResponseBindStruct{})
	fmt.Printf("tSlice: %#v\n", tSlice)

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

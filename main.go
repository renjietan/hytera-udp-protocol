package main

import (
	"fmt"
	"time"
)

func main() {
	c, err := NewUdpClient("8.135.10.183:55022")
	if err != nil {
		c.OnMsg = func(data string, err error) {
			fmt.Println(data, err)
		}
		c.OnClose = func(data string, err error) {
			fmt.Println(data)
		}
	}
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//// 关闭应用程序
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := c.; err != nil {
	//
	//}
	fmt.Println("c=================", c)
	fmt.Println("err=================", err)
	time.Sleep(100000 * time.Second)
}

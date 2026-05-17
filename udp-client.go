package main

import (
	"fmt"
	"net"
	"time"

	"github.com/renjietan/hytera-udp-protocol/options"
	"github.com/renjietan/hytera-udp-protocol/tools"
)

type UdpClient struct {
	conn           *net.UDPConn
	done           chan struct{}
	options        *options.App
	timeoutManager *tools.TimeoutManager
}

func NewUdpClient(addr *options.App) (*UdpClient, error) {
	//localaddr := fmt.Sprintf("%s:%s", addr.Host, addr.Port)
	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		return nil, fmt.Errorf("不是有效的udp地址: %v", err.Error())
	}
	conn, connErr := net.ListenUDP("udp", udpAddr)
	if connErr != nil {
		return nil, fmt.Errorf("连接失败: %v", connErr.Error())
	}

	client := &UdpClient{
		conn:           conn,
		options:        addr,
		timeoutManager: tools.NewTimeoutManager(),
	}
	err = client.Send("hello")
	if err != nil {
		return nil, err
	}
	go client.onMsg()
	return client, nil
}

func (c *UdpClient) onMsg() {
	buf := make([]byte, 1024)
	for {
		_ = c.conn.SetReadDeadline(time.Now().Add(2 * time.Second))

		n, addr, err := c.conn.ReadFromUDP(buf)
		if err != nil {
			// TODO：读取超时等错误是否需要处理，暂不清楚，此处不做处理，继续循环
			netErr, ok := err.(net.Error)
			if ok && netErr.Timeout() {
				continue
			}
			fmt.Println("接收出错:", err)
			continue
		}

		msg := buf[:n]
		//c.mu.Lock()
		//c.mu.Unlock()
		if c.options.OnMsgFunc != nil {
			c.options.OnMsgFunc(addr, string(msg), nil)
		}
		c.timeoutManager.Set("ping", 3*time.Second, func() {
			err := c.Send("ping")
			if err != nil {
				c.options.OnCloseFunc(addr, "心跳发送失败", err)
			}
		})
	}
}

func (c *UdpClient) onClose() {
	select {
	case <-c.done:
		fmt.Println("用于消息接收的协程退出")
		return
	default:
	}
}

//	func (c *UdpClient) onError(msg interface{}) {
//		if c.options.OnErrorFunc != nil {
//
//		}
//	}
func (c *UdpClient) Send(str string) error {
	remoteAddr := fmt.Sprintf("%s:%s", c.options.RHost, c.options.RPort)
	remote, _ := net.ResolveUDPAddr("udp", remoteAddr)
	_, err := c.conn.WriteToUDP([]byte(str), remote)
	if err != nil {
		return err
	}
	return nil
}

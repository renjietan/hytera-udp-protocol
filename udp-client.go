package main

import (
	"fmt"
	"net"
	"time"
)

type Callback func(data string, err error)

type UdpClient struct {
	conn    *net.UDPConn
	done    chan struct{}
	OnMsg   Callback
	OnClose Callback
	//mu      sync.RWMutex
}

func NewUdpClient(addr string) (*UdpClient, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("不是有效的udp地址: %v", err.Error())
	}
	conn, conn_err := net.DialUDP("udp", nil, udpAddr)
	if conn_err != nil {
		return nil, fmt.Errorf("连接失败: %v", conn_err.Error())
	}
	client := &UdpClient{
		conn: conn,
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
		if c.OnMsg != nil {
			c.OnMsg(string(msg), nil)
		}
		fmt.Printf("收到来自 %s 的消息: %+v\n", addr.String(), msg)
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

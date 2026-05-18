package main

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/options"
	"github.com/renjietan/hytera-udp-protocol/tools"
)

type UdpClient struct {
	conn           *net.UDPConn
	done           chan struct{}
	options        *options.App
	timeoutManager *tools.TimeoutManager
	rAddress       *net.UDPAddr
}

func NewUdpClient(config *options.App) (*UdpClient, error) {
	// 校验: 本机IP
	address := fmt.Sprintf(":%s", config.Port)
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, fmt.Errorf("不是有效的udp地址: %v", err.Error())
	}
	// 校验: 远程IP
	address = fmt.Sprintf("%s:%s", config.RHost, config.RPort)
	rAddress, rErr := net.ResolveUDPAddr("udp", address)
	if rErr != nil {
		return nil, fmt.Errorf("不是有效的远程地址: %v", rErr.Error())
	}
	// 启动UDP
	conn, connErr := net.ListenUDP("udp", udpAddr)
	if connErr != nil {
		return nil, fmt.Errorf("连接失败: %v", connErr.Error())
	}

	client := &UdpClient{
		conn:           conn,
		options:        config,
		timeoutManager: tools.NewTimeoutManager(),
		rAddress:       rAddress,
	}
	defer func(v *net.UDPAddr) {
		b := core.ResultLogin("admin", 11, true)
		if b == nil {
			fmt.Println("错误的字节")
			return
		}
		e := client.Send(b)
		if e != nil {
			client.options.OnErrorFunc(v, nil, e)
			return
		}
		client.timeoutManager.Set("login", 3*time.Second, func() {
			client.options.OnErrorFunc(v, nil, errors.New("登录超时"))
		})
	}(rAddress)
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
			var netErr net.Error
			ok := errors.As(err, &netErr)
			if ok && netErr.Timeout() {
				continue
			}
			fmt.Println("onMsg 读取超时:", err)
			continue
		}
		msg := buf[:n]
		if c.options.OnMsgFunc != nil {
			c.options.OnMsgFunc(addr, msg, nil)
		}
		//c.timeoutManager.Set("ping", 3*time.Second, func() {
		//	c.Send()
		//})
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
func (c *UdpClient) Send(b []byte) error {
	remoteAddr := fmt.Sprintf("%s:%s", c.options.RHost, c.options.RPort)
	remote, _ := net.ResolveUDPAddr("udp", remoteAddr)
	_, err := c.conn.WriteToUDP(b, remote)
	if err != nil {
		return err
	}
	return nil
}

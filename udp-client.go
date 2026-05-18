package main

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/options"
	"github.com/renjietan/hytera-udp-protocol/options/enums"
	"github.com/renjietan/hytera-udp-protocol/tools"
)

type UdpClient struct {
	conn           *net.UDPConn
	done           chan struct{}
	options        *options.App
	timeoutManager *tools.TimeoutManager
}

func NewUdpClient(config *options.App) (*UdpClient, error) {
	// 校验: 本机IP
	address := fmt.Sprintf(":%s", config.Port)
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, fmt.Errorf("不是有效的udp地址: %v", err.Error())
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
			var netErr net.Error
			ok := errors.As(err, &netErr)
			if ok && netErr.Timeout() {
				continue
			}
			c.options.OnErrorFunc(options.Error("", err.Error(), addr))
			continue
		}
		data := buf[:n]
		if c.options.OnMsgFunc != nil {
			c.options.OnMsgFunc(options.Success("", data, "", addr))
		}
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
func (c *UdpClient) Send(RHost string, RPort int, event string, b []byte) {
	remoteAddr, err := tools.GenerateAddress(RHost, RPort)
	if err != nil {
		c.options.OnErrorFunc(options.Error("", enums.InvalidAddress, nil))
		return
	}
	_, err = c.conn.WriteToUDP(b, remoteAddr)
	if err != nil {
		c.options.OnErrorFunc(options.Error(event, err.Error(), remoteAddr))
		return
	}
	c.timeoutManager.Set(event, c.options.Duration, func() {
		// TODO: 收到回复后 删除 定时器
		c.options.OnErrorFunc(options.Error(event, enums.ReplyTimeout, remoteAddr))
	})
}

func (c *UdpClient) Login(RHost string, RPort int, username string, userId int, alive bool) {
	loginByte, _ := core.TempLogin(username, userId, alive)
	recordField := tools.GetRecursiveField(loginByte, []options.Item{})
	_, res := tools.Struct2Bytes(loginByte, recordField, []byte{})
	c.Send(RHost, RPort, "login", res)
}

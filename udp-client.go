package main

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/renjietan/hytera-udp-protocol/core/request/auth"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
	"github.com/renjietan/hytera-udp-protocol/types/enums"
)

type UdpClient struct {
	conn           *net.UDPConn
	done           chan struct{}
	options        *types.App
	timeoutManager *tools.TimeoutManager
}

func NewUdpClient(config *types.App) (*UdpClient, error) {
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
	if config.Duration <= 0 {
		config.Duration = 3 * time.Second
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
			var netErr net.Error
			ok := errors.As(err, &netErr)
			if ok && netErr.Timeout() {
				continue
			}
			c.options.OnErrorFunc(tools.Error(enums.EventReader, err.Error(), addr))
			continue
		}
		data := buf[:n]
		Biz := tools.NewSafeBytes(data).Slice(8, enums.Backward)
		SAP := Biz.GetByte(0)
		if SAP == 1 { // 登录 开启心跳
			c.AuthReply(data, addr)
		}
		if c.options.OnMsgFunc != nil {
			c.options.OnMsgFunc(tools.Success("", data, "", addr))
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
//		if c.types.OnErrorFunc != nil {
//
//		}
//	}
func (c *UdpClient) Send(RHost string, RPort int, event string, b []byte) {
	remoteAddr, err := tools.GenerateAddress(RHost, RPort)
	if err != nil {
		c.options.OnErrorFunc(tools.Error("", enums.InvalidAddress, nil))
		return
	}
	_, err = c.conn.WriteToUDP(b, remoteAddr)
	if err != nil {
		c.options.OnErrorFunc(tools.Error(event, err.Error(), remoteAddr))
		return
	}
	c.timeoutManager.Set(event, c.options.Duration, func() {
		c.options.OnErrorFunc(tools.Error(event, enums.ReplyTimeout, remoteAddr))
	})
}

func (c *UdpClient) Login(RHost string, RPort int, username string, userId int, alive bool) {
	address, errAddr := tools.GenerateAddress(RHost, RPort)
	if errAddr != nil {
		c.options.OnErrorFunc(tools.Error(enums.EventLogin, enums.InvalidAddress, nil))
	}
	res, err := core_request_auth.LoginReq(username, userId, int(c.options.Duration.Milliseconds()))
	if err != nil {
		c.options.OnErrorFunc(tools.Error(enums.EventLogin, err.Error(), address))
		return
	}
	fmt.Printf("login byte:%#v\n", res)
	c.Send(RHost, RPort, enums.EventLogin, res)
}

func (c *UdpClient) Ping(RHost string, RPort int, userId int) {
	address, errAddr := tools.GenerateAddress(RHost, RPort)
	if errAddr != nil {
		c.options.OnErrorFunc(tools.Error(enums.EventPing, enums.InvalidAddress, nil))
		return
	}
	res, err := core_request_auth.SuperviseReq(userId)
	if err != nil {
		c.options.OnErrorFunc(tools.Error(enums.EventPing, err.Error(), address))
		return
	}
	c.Send(RHost, RPort, enums.EventPing, res)
}

func (c *UdpClient) Logout(RHost string, RPort int, userId int) {

}

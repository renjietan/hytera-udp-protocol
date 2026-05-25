package client

import (
	"net"

	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types/enums"
)

// AuthReply 回复
func (c *UdpClient) AuthReply(data []byte, address *net.UDPAddr) {
	_data := tools.NewSafeBytes(data)
	UserID := _data.GetByte(7)
	OptCode := _data.GetByte(9)
	if OptCode == 0x85 { // 133-登录
		c.timeoutManager.StopAll()
		c.Ping(address.IP.String(), address.Port, int(UserID))
		c.timeoutManager.Set(enums.EventPing, c.options.Duration, func() {
			c.options.OnErrorFunc(Error(enums.EventPing, enums.ReplyTimeout, address))
		})
	} else if OptCode == 0x83 { // 131-心跳
		status := _data.GetByte(10)
		if status == 0x01 {
			c.timeoutManager.StopAll()
			c.options.OnErrorFunc(Error(enums.EventPing, enums.PingDisconnect, address))
			return
		}
		c.timeoutManager.StopByName(enums.EventPing)
		c.timeoutManager.Set(enums.EventPing, c.options.Duration, func() {
			c.Ping(address.IP.String(), address.Port, int(UserID))
			c.timeoutManager.Set(enums.EventPing, c.options.Duration, func() {
				c.options.OnErrorFunc(Error(enums.EventPing, enums.ReplyTimeout, address))
			})
		})
	}
}

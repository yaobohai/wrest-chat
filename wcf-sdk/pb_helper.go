package wcf

import (
	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/v3/protocol"
	"go.nanomsg.org/mangos/v3/protocol/pair1"
	"go.nanomsg.org/mangos/v3/transport/all"
	"google.golang.org/protobuf/proto"

	"github.com/opentdp/go-helper/logman"
)

// 通用消息

type cmdMsg struct {
	*Request
}

// 生成消息
func (c *cmdMsg) build() []byte {
	marshal, _ := proto.Marshal(c)
	return marshal
}

// 生成请求消息
func genFunReq(fun Functions) *cmdMsg {
	return &cmdMsg{
		&Request{Func: fun, Msg: nil},
	}
}

// RPC 客户端

type pbSocket struct {
	Server string // 接口地址
	socket protocol.Socket
}

// 连接服务器
func (c *pbSocket) dial() (err error) {
	all.AddTransports(nil) // 注册所有传输协议
	logman.Info("pbsocket", "server", c.Server)
	c.socket, err = pair1.NewSocket()
	if err != nil {
		return err
	}
	return c.socket.Dial(c.Server)
}

// 接收数据
func (c *pbSocket) recv() (*Response, error) {
	resp := &Response{}
	recv, err := c.socket.Recv()
	if err == nil {
		err = proto.Unmarshal(recv, resp)
	}
	return resp, err
}

// 发送数据
func (c *pbSocket) send(data []byte) error {
	return c.socket.Send(data)
}

// 关闭连接
func (c *pbSocket) close() error {
	return c.socket.Close()
}

// 设置超时时间
func (c *pbSocket) deadline(d int) {
	c.socket.SetOption(mangos.OptionRecvDeadline, d)
	c.socket.SetOption(mangos.OptionSendDeadline, d)
}

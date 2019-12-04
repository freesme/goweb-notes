package main

import (
	"log"
	"net"
	"net/rpc"
)

type Params struct {
	Width, Height int
}
type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main() {
	rect := new(Rect)
	//注册rect服务
	rpc.Register(rect)
	//服务端口
	//在此处我们采用了TCP协议，然后需要自己控制连接，当有客户端连接上来后，我们需要把这个连接交给rpc来处理
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":2001")
	if err != nil {
		log.Fatal(err)
	}

	//监听端口
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//手动控制rpc
		rpc.ServeConn(conn)
	}
}

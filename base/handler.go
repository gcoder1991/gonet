package base

import "net"

type TcpHandler interface {
	OnActive(conn *net.TCPConn)
	OnInactive(conn *net.TCPConn)
	OnRead(conn *net.TCPConn, protocol Protocol)
	OnError(conn *net.TCPConn, err error)
}

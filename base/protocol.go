package base

import "net"

type Protocol struct {
	Parser func(conn *net.Conn) (p Protocol)
}

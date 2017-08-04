package base

import (
	"net"
)

type Protocol struct {
	ProtocolProcessor
}

type ProtocolProcessor interface {
	Processor(protocol Protocol) error
	Parser(conn *net.TCPConn) (p Protocol, err error)
}


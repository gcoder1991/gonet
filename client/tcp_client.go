package client

import (
	"github.com/gcoder1991/gonet/base"
	"net"
	"sync"
)

type TcpClient struct {
	Addr *net.TCPAddr
	base.TcpHandler

	protocol base.Protocol

	closeWait *sync.WaitGroup
}

func (tc TcpClient) Connect() error {
	tcpConn, err := net.DialTCP("tcp", nil, tc.Addr)
	if err != nil {
		return err
	}
	defer func() {
		tcpConn.Close()
		tc.OnInactive(tcpConn)
	}()
	for {
		tc.OnActive(tcpConn)
		for {
			p, err := tc.protocol.Parser(tcpConn)
			if err != nil {
				tc.OnError(tcpConn, err)
				continue
			}
			tc.OnRead(tcpConn, p)
		}
	}
}

package server

import (
	"fmt"
	"github.com/gcoder1991/gonet/base"
	"log"
	"net"
	"sync"
)

type TcpServer struct {
	protocol base.Protocol

	Addr     *net.TCPAddr
	Listener *net.TCPListener

	Handler

	closeWait *sync.WaitGroup
}

type Handler interface {
	OnActive(conn *net.TCPConn)
	OnInactive(conn *net.TCPConn)
	OnRead(conn *net.TCPConn, protocol base.Protocol)
	OnError(conn *net.TCPConn, err error)
}

func (ts *TcpServer) Start() error {
	listener, e := net.ListenTCP("tcp", ts.Addr)
	defer listener.Close()
	if e != nil {
		return e
	}

	for {
		tcpConn, e := listener.AcceptTCP()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Fatal(fmt.Sprintf("tcp: Accept error: %v", e))
				continue
			}
			return e
		}

		go func() {
			defer func() {
				tcpConn.Close()
				ts.OnInactive(tcpConn)
			}()
			ts.OnActive(tcpConn)
			for {
				p, err := ts.protocol.Parser(tcpConn)
				if err != nil {
					ts.OnError(tcpConn, err)
					continue
				}
				ts.OnRead(tcpConn, p)
			}
		}()
	}
}

func (ts *TcpServer) Stop() {
	ts.Listener.Close()
	ts.closeWait.Done()
}

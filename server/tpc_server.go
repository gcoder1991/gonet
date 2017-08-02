package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type TcpServer struct {
	Addr     *net.TCPAddr
	Listener *net.TCPListener

	closeWait *sync.WaitGroup

	OnActive   func(conn *net.TCPConn)
	OnInactive func(conn *net.TCPConn)
	OnRead     func(conn *net.TCPConn)
	OnError    func(conn *net.TCPConn)
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
			go ts.OnActive(tcpConn)
			// TODO
		}()
	}
}

func (ts *TcpServer) Stop() {
	ts.Listener.Close()
	ts.closeWait.Done()
}

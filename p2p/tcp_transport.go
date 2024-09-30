package p2p

import (
	"fmt"
	"net"
	"sync"
)

// this struct represent the remote node over a TCP established connection
type TCPPeer struct {
	conn net.Conn
	// if we dial and retrive a con => outbound == true
	// if we accept and retrive a conn => outbound == false

	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	mu            sync.RWMutex
	shakeHands    Handshakefunc
	decoder       Decoder
	peers         map[net.Addr]Peer
}

// method for TCPTransport structure
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    NOPHandshakeFunc,
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		fmt.Printf("new incoming connection %+v\n", conn)
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.shakeHands(peer); err != nil {

	}

	// read loop
	 Temp := 
	 msg := &Temp{}
	 
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("tcp error : %s\n", err)
			continue
		}
	}
}

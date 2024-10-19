package p2p

import (
	"fmt"
	"log"
	"net"
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

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

// various opts to configure the TCP transport
type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
	OnPeer        func(Peer) error
}

// responsible for listening for incoming connections , accepting them , handling requests
type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcch    chan RPC
}

// method for TCPTransport structure
func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

// Consume implements the transport interface
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

// Listens on specific address and port for incoming connections
// and hands off these connections to startAcceptLoop for processing
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	log.Printf("TCP transport listening on port: %s\n", t.ListenAddr)

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		go t.handleConn(conn)
	}
}

type Temp struct{}

// this is were actual p2p coomn happens , after a connection is accepted , a TCPPeer is created
// and a handshake is performed. after handshake node inters a loop and it reads messages(RPCS)
// from the peer.
func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error
	defer func() {
		fmt.Printf("dropping peer connection: %s", err)
		conn.Close()
	}()
	peer := NewTCPPeer(conn, true)

	if err = t.HandshakeFunc(peer); err != nil {
		return
	}

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}
	// read loop
	rpc := RPC{}
	for {
		err = t.Decoder.Decode(conn, &rpc)
		if err != nil {
			return
		}
		rpc.From = conn.RemoteAddr()
		t.rpcch <- rpc
	}
}

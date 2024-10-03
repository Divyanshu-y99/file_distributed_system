package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPtransport(t *testing.T) {
	opts := TCPTransportOpts{
		ListenAddr:    ":3030",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTCPTransport(opts)
	assert.Equal(t, tr.ListenAddr, ":3030") // compare both the addresses and provide the result in bool form

	assert.Nil(t, tr.ListenAndAccept()) // if the object is nill
}

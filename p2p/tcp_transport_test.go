package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPtransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)
	assert.Equal(t, tr.listenAddress, listenAddr) // compare both the addresses and provide the result in bool form

	assert.Nil(t, tr.ListenAndAccept()) // if the object is nill
}

package p2p

// Peer is the interface that represents the remote node
type Peer interface{}

// Transport is any thing that handles the communication
// between the nodes in the network. This can ibe of the
// form (TCP,UDP...)

type Transport interface {
	ListenAndAccept() error
}

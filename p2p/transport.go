package p2p

// Peer is the interface that represents the remote node
type Peer interface {
	Close() error
}

// Transport is any thing that handles the communication
// between the nodes in the network. This can be of the
// form (TCP,UDP...)

type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}

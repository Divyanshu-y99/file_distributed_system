package p2p

type Handshakefunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }

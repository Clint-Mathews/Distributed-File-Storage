package p2p

// Peer is an interface that represents the remote node.
type Peer interface{}

// Transport is anything that handles commuincation between the nodes in network.
// This can be of the form (TCP, UDP, Websockets, ...)
type Transport interface {
	ListenAndAccept() error
}

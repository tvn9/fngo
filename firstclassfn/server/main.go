// Server constructor
package main

type (
	ServerOptions func(options) options
	TransportType int
)

const (
	UDP TransportType = iota
	TCP
)

type Server struct {
	options
}

type options struct {
	MaxConnection int
	TransportType TransportType
	Name          string
}

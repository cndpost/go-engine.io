package websocket

import (
	"github.com/cndpost/go-engine.io/v1.0/transport"
)

var Creater = transport.Creater{
	Name:      "websocket",
	Upgrading: true,
	Server:    NewServer,
	Client:    NewClient,
}

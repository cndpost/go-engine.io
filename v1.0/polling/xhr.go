package polling

import (
	"github.com/cndpost/go-engine.io/v1.0/transport"
)

var Creater = transport.Creater{
	Name:      "polling",
	Upgrading: false,
	Server:    NewServer,
	Client:    NewClient,
}

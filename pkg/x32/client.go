package x32

import (
	"github.com/jonavdm/x32/pkg/osc"
)

type Client struct {
	Connection *osc.Client
}

func NewClient(address string) (Client, error) {
	connection := osc.NewClient(address)
	err := connection.Connect()
	return Client{&connection}, err
}

func (c *Client) Close() {
	c.Connection.Close()
}

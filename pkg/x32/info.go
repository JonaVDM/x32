package x32

import "github.com/jonavdm/x32/pkg/osc"

func (c *Client) SendExpansionInfo() {
	c.Connection.Send(osc.Message{Message: "/‐stat/xcardtype"})
}

func (c *Client) GetXInfo() (*osc.Message, error) {
	return c.Connection.GetCommand(osc.Message{Message: "/xinfo"})
}

func (c *Client) GetStatus() (*osc.Message, error) {
	return c.Connection.GetCommand(osc.Message{Message: "/status"})
}

package x32

import (
	"errors"
	"net"
	"time"
)

func NewClient(address string) Client {
	return Client{
		Address:       address,
		Message:       make(chan Message),
		StopHeartBeat: make(chan bool),
		StopSend:      make(chan bool),
	}
}

type Client struct {
	Address       string
	Connection    net.Conn
	Message       chan Message
	StopHeartBeat chan bool
	StopSend      chan bool
}

func (c *Client) Connect() error {
	if c.Connection != nil {
		return errors.New("connection already opened")
	}

	if _, err := net.ResolveUDPAddr("udp", c.Address); err != nil {
		return err
	}

	var err error
	c.Connection, err = net.Dial("udp", c.Address)
	if err != nil {
		return err
	}

	go c.SendLoop()
	go c.Heartbeat()

	return nil
}

func (c *Client) Close() {
	if c.Connection != nil {
		c.StopHeartBeat <- true
		c.StopSend <- true
		c.Connection.Close()
	}
}

func (c *Client) Heartbeat() {
	ticker := time.NewTicker(time.Second * 8)
	c.Message <- Message{Message: MessageXRemote}
	for {
		select {
		case <-c.StopHeartBeat:
			return

		case <-ticker.C:
			c.Message <- Message{Message: MessageXRemote}
		}
	}
}

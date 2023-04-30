package x32

import (
	"errors"
	"net"
	"time"
)

func NewClient(address string) Client {
	return Client{
		address:       address,
		message:       make(chan Message),
		stopHeartBeat: make(chan bool),
		stopSend:      make(chan bool),
	}
}

type Client struct {
	address       string
	connection    net.Conn
	message       chan Message
	stopHeartBeat chan bool
	stopSend      chan bool
}

func (c *Client) Connect() error {
	if c.connection != nil {
		return errors.New("connection already opened")
	}

	if _, err := net.ResolveUDPAddr("udp", c.address); err != nil {
		return err
	}

	var err error
	c.connection, err = net.Dial("udp", c.address)
	if err != nil {
		return err
	}

	go c.sendLoop()
	go c.heartbeat()

	return nil
}

func (c *Client) Close() {
	if c.connection != nil {
		c.stopHeartBeat <- true
		c.stopSend <- true
		c.connection.Close()
	}
}

func (c *Client) heartbeat() {
	ticker := time.NewTicker(time.Second * 9)
	c.message <- Message{Message: MessageXRemote}
	for {
		select {
		case <-c.stopHeartBeat:
			return

		case <-ticker.C:
			c.message <- Message{Message: MessageXRemote}
		}
	}
}

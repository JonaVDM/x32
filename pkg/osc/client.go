package osc

import (
	"errors"
	"net"
	"time"
)

func NewClient(address string) Client {
	return Client{
		address:       address,
		Message:       make(chan Message),
		stopHeartBeat: make(chan bool),
		stopSend:      make(chan bool),
	}
}

type Client struct {
	address       string
	connection    net.Conn
	Message       chan Message
	stopHeartBeat chan bool
	stopSend      chan bool
	subscriptions []chan Message
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
	go c.listen()

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
	c.Message <- Message{Message: MessageXRemote}
	for {
		select {
		case <-c.stopHeartBeat:
			return

		case <-ticker.C:
			c.Message <- Message{Message: MessageXRemote}
		}
	}
}

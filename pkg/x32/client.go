package x32

import (
	"bytes"
	"errors"
	"net"
	"time"
)

func NewClient(address string) Client {
	return Client{
		Address:       address,
		Message:       make(chan string),
		StopHeartBeat: make(chan bool),
		StopSend:      make(chan bool),
	}
}

type Client struct {
	Address       string
	Connection    net.Conn
	Message       chan string
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

func (c *Client) SendLoop() {
	for {
		select {
		case message := <-c.Message:
			code := []byte(message)
			code = bytes.ReplaceAll(code, []byte("~"), []byte{0x0})
			_, err := c.Connection.Write(code)
			if err != nil {
				panic(err)
			}

		case <-c.StopSend:
			return
		}
	}
}

func (c *Client) Heartbeat() {
	ticker := time.NewTicker(time.Second * 8)
	c.Message <- "/xremote"
	for {
		select {
		case <-c.StopHeartBeat:
			return

		case <-ticker.C:
			c.Message <- "/xremote"
		}
	}
}

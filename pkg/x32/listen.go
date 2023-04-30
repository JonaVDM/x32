package x32

import (
	"bytes"
)

func (c *Client) Subscribe(subscription chan Message) {
	c.subscriptions = append(c.subscriptions, subscription)
}

func (c *Client) UnSubscribe(subscription chan Message) {
	for i, l := range c.subscriptions {
		if l == subscription {
			c.subscriptions = append(c.subscriptions[:i], c.subscriptions[i+1:]...)
			break
		}
	}
}

func (c *Client) callSubscribers(message Message) {
	for _, listener := range c.subscriptions {
		go func(ch chan Message) {
			ch <- message
		}(listener)
	}
}

func (c *Client) listen() {
	for {
		buffer := make([]byte, 1024)
		size, err := c.connection.Read(buffer)

		if err != nil {
			panic(err)
		}

		// ignore empty messages
		if size == 0 {
			continue
		}

		nullIndex := bytes.Index(buffer, []byte{0})
		msg := Message{
			Message: string(buffer[:nullIndex]),
		}

		index := bytes.Index(buffer, []byte{44})
		nullIndex = bytes.Index(buffer[index:], []byte{0})

		msg.Values = make([]Value, nullIndex-1)

		block := index + ((nullIndex)/4+1)*4
		for i, parameterType := range buffer[index+1 : index+nullIndex] {
			msg.Values[i] = Value{
				Type:  parameterType,
				Value: buffer[block : block+4],
			}

			if parameterType == 's' && msg.Values[i].Value[3] != 0 {
				for {
					block += 4
					ni := bytes.Index(buffer[block:block+4], []byte{0})

					if ni != -1 {
						msg.Values[i].Value = append(msg.Values[i].Value, buffer[block:block+ni]...)
						break
					}

					msg.Values[i].Value = append(msg.Values[i].Value, buffer[block:block+4]...)
				}
			}

			block += 4
		}

		c.callSubscribers(msg)
	}
}

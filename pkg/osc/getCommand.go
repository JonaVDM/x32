package osc

import (
	"errors"
	"time"
)

func (c *Client) GetCommand(message Message) (*Message, error) {
	ch := make(chan Message)
	c.Subscribe(ch)
	c.Send(message)

	timeout := time.NewTimer(time.Second * 3)

	for {
		select {
		case <-timeout.C:
			c.UnSubscribe(ch)
			close(ch)

			return nil, errors.New("request timeout")

		case inc, ok := <-ch:
			if !ok {
				return nil, errors.New("error in connection")
			}

			if inc.Message != message.Message {
				continue
			}

			timeout.Stop()
			c.UnSubscribe(ch)
			close(ch)

			return &inc, nil
		}
	}
}

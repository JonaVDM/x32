package x32

import (
	"fmt"

	"github.com/jonavdm/x32/internal/utils"
)

func (c *Client) SetMixbusFader(channel, mixbus int, value float32) error {
	bytes, err := utils.Float32ToBytes(value)
	if err != nil {
		return err
	}

	c.Message <- Message{
		Message: fmt.Sprintf("/ch/%02d/mix/%02d/level", channel, mixbus),
		Values: []Value{
			{
				Type:  'f',
				Value: bytes,
			},
		},
	}

	return nil
}

func (c *Client) SetFader(channel int, value float32) error {
	bytes, err := utils.Float32ToBytes(value)
	if err != nil {
		return err
	}

	c.Message <- Message{
		Message: fmt.Sprintf("/ch/%02d/mix/fader", channel),
		Values: []Value{
			{
				Type:  'f',
				Value: bytes,
			},
		},
	}
	return nil
}

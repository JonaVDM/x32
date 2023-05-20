package x32

import (
	"encoding/binary"
	"fmt"

	"github.com/jonavdm/x32/internal/utils"
	"github.com/jonavdm/x32/pkg/osc"
)

func (c *Client) SetMixbusFader(channel, mixbus int, value float32) error {
	bytes, err := utils.Float32ToBytes(value)
	if err != nil {
		return err
	}

	c.Connection.Message <- osc.Message{
		Message: fmt.Sprintf("/ch/%02d/mix/%02d/level", channel, mixbus),
		Values: []osc.Value{
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

	c.Connection.Message <- osc.Message{
		Message: fmt.Sprintf("/ch/%02d/mix/fader", channel),
		Values: []osc.Value{
			{
				Type:  'f',
				Value: bytes,
			},
		},
	}
	return nil
}

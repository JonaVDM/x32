package x32

import (
	"fmt"

	"github.com/jonavdm/x32/internal/convert"
)

func (c *Client) SetMixbusFader(channel, mixbus int, value float32) error {
	bytes, err := convert.Float32ToBytes(value)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("/ch/%02d/mix/%02d/level~,f~~%s", channel, mixbus, string(bytes))
	c.Message <- message
	return nil
}

func (c *Client) SetFader(channel int, value float32) error {
	bytes, err := convert.Float32ToBytes(value)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("/ch/%02d/mix/fader~~~~,f~~%s", channel, string(bytes))
	c.Message <- message
	return nil
}

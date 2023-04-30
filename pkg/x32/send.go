package x32

import (
	"github.com/jonavdm/x32/internal/utils"
)

func (c *Client) Send(msg Message) {
	c.message <- msg
}

func (c *Client) sendLoop() {
	for {
		select {
		case message := <-c.message:
			code := utils.PadBytes([]byte(message.Message))

			if len(message.Values) > 0 {
				types := make([]byte, len(message.Values)+1)
				types[0] = ','
				values := make([]byte, 0)
				for i, value := range message.Values {
					types[i+1] = value.Type
					values = append(values, value.Value...)
				}

				code = append(code, utils.PadBytes(types)...)
				code = append(code, values...)
			}

			_, err := c.connection.Write(code)
			if err != nil {
				panic(err)
			}

		case <-c.stopSend:
			return
		}
	}
}

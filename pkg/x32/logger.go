package x32

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (c *Client) Logger() {
	ch := make(chan Message)
	c.Subscribe(ch)
	defer func() {
		c.UnSubscribe(ch)
		close(ch)
	}()

	for {
		msg := <-ch

		fmt.Printf("%s", msg.Message)
		for _, value := range msg.Values {
			fmt.Printf(" [%c %s]", value.Type, valueToText(value.Type, value.Value))
		}
		fmt.Println("")
	}
}

func valueToText(paramType byte, value []byte) string {
	switch paramType {
	case 's':
		return string(value)

	case 'i':
		buffer := bytes.NewReader(value)
		var num uint32
		binary.Read(buffer, binary.BigEndian, &num)
		return fmt.Sprintf("%d", num)

	case 'f':
		buffer := bytes.NewReader(value)
		var num float32
		binary.Read(buffer, binary.BigEndian, &num)
		return fmt.Sprintf("%f", num)

	default:
		return ""
	}
}

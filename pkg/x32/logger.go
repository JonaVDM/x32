package x32

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (c *Client) Logger() {
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

		index := bytes.Index(buffer, []byte{44})
		fmt.Printf("%s %c [", buffer[:index], buffer[index+1])
		fmt.Print(valueToText(buffer[index+1 : size]))
		fmt.Println("]")
	}
}

func valueToText(value []byte) string {
	switch value[0] {
	case 's':
		return string(value[3:])

	case 'i':
		buffer := bytes.NewReader(value[3:])
		var num uint32
		binary.Read(buffer, binary.BigEndian, &num)
		return fmt.Sprintf("%d", num)

	case 'f':
		buffer := bytes.NewReader(value[3:])
		var num float32
		binary.Read(buffer, binary.BigEndian, &num)
		return fmt.Sprintf("%f", num)

	default:
		out := ""
		for _, b := range value[2:] {
			out += fmt.Sprintf("%x ", b)
		}

		return out
	}
}

package convert

import (
	"bytes"
	"encoding/binary"
)

func Float32ToBytes(num float32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, num)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

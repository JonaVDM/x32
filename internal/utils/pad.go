package utils

// PadBytes will pad the value with up to 4 bytes, adhering to OSC
func PadBytes(value []byte) []byte {
	add := len(value) % 4

	return append(value, make([]byte, 4-add)...)
}

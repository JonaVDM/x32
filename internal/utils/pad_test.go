package utils_test

import (
	"fmt"
	"testing"

	"github.com/jonavdm/x32/internal/utils"
	"gotest.tools/v3/assert"
)

func TestPadBytes(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "With 1 byte",
			input:    []byte{10},
			expected: []byte{10, 0, 0, 0},
		},
		{
			name:     "With 4 bytes",
			input:    []byte{10, 92, 3, 50},
			expected: []byte{10, 92, 3, 50, 0, 0, 0, 0},
		},
		{
			name:     "With 3 byte",
			input:    []byte{10, 20, 30},
			expected: []byte{10, 20, 30, 0},
		},
		{
			name:     "With 5 byte",
			input:    []byte{5, 1, 40, 8, 15},
			expected: []byte{5, 1, 40, 8, 15, 0, 0, 0},
		},
		{
			name:     "With 8 byte",
			input:    []byte{6, 10, 21, 50, 40, 33, 70, 20},
			expected: []byte{6, 10, 21, 50, 40, 33, 70, 20, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		fmt.Println(utils.PadBytes(tc.input))
		t.Run(tc.name, func(t *testing.T) {
			assert.DeepEqual(t, utils.PadBytes(tc.input), tc.expected)
		})
	}
}

package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMD5Hash(t *testing.T) {
	tests := []struct { // добавляем слайс тестов
		name   string
		input  []byte
		result string
	}{
		{
			name:   "simple test",
			input:  []byte("ivan"),
			result: "2c42e5cf1cdbafea04ed267018ef1511",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, GetMD5Hash(test.input), test.result)
		})
	}
}

func Example() {
	GetMD5Hash([]byte("simple test"))

	// Output
	// 2c42e5cf1cdbafea04ed267018ef1511
}

package helper

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// test from struct array
func TestTable(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "HelloWorld(Ryan)",
			request: "Ryan",
			expected: "Hello, Ryan",
		},
		{
			name: "HelloWorld(Adit)",
			request: "Adit",
			expected: "Hello, Adit",
		},
		{
			name: "HelloWorld(Putra)",
			request: "Putra",
			expected: "Hello, Putra",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

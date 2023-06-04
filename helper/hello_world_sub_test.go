package helper

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// test subtest
func TestSubtest(t *testing.T) {
	t.Run("Ryan", func(t *testing.T) {
		result := HelloWorld("Ryan")
		assert.Equal(t, "Hello, Ryan", result, "Result is not Hello, Ryan")
	})
	t.Run("Adit", func(t *testing.T) {
		result := HelloWorld("Adit")
		assert.Equal(t, "Hello, Adit", result, "Result is not Hello, Adit")
	})
}
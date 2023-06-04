package helper

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldAssert(t *testing.T) {
	// assert refer to Fail, require refer to FailNow
	result := HelloWorld("Adit")
	assert.Equal(t, "Hello, Adit", result, "Result must be 'Hello, Adit'")
}
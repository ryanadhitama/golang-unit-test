package helper

import (
	"runtime"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldSkip(t *testing.T) {
	// go test -v -run=TestHelloWorldSkip
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run at macOS")
	}
	result := HelloWorld("Adit")
	assert.Equal(t, "Hello, Adit", result, "Result must be 'Hello, Adit'")
}
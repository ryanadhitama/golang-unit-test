package helper

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Test")

	m.Run()

	fmt.Println("After Test")
}

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

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ryan")

	if result != "Hello, Ryan" {
		// t.Fail()
		// t.FailNow()
		t.Error("Result is not Hello, Ryan")
		// t.Fatal("Result is not Hello, Ryan")
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldTrue(t *testing.T) {
	result := HelloWorld("Ryan")

	if result != "Hello, Ryan" {
		t.Fail()
	}
}

func TestHelloWorldAssert(t *testing.T) {
	// assert refer to Fail, require refer to FailNow
	result := HelloWorld("Adit")
	assert.Equal(t, "Hello, Adit", result, "Result must be 'Hello, Adit'")
}
package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ryan")

	if result != "Hello, Ryan" {
		panic("Result is not Hello, Ryan")
	}
}
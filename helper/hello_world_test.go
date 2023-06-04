package helper

import (
	"fmt"
	"testing"
)

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
package helper

import (
	"fmt"
	"testing"
)

// before and after test
func TestMain(m *testing.M) {
	fmt.Println("Before Test")

	m.Run()

	fmt.Println("After Test")
}
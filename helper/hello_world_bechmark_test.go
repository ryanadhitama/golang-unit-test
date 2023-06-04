package helper

import (
	"testing"
)

// running go test -v ./... -run=NotUnitTest -bench=.

func BenchmarkHelloAdit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Adit")
	}
}

func BenchmarkHelloAdhitama(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Adhitama")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Ryan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ryan")
		}
	})

	b.Run("Adit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Adit")
		}
	})
}
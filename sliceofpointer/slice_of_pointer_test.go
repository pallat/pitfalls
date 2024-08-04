package sliceofpointer

import (
	"testing"
)

func BenchmarkSliceOfPointer(b *testing.B) {
	for range b.N {
		DoWithPointer()
	}
}

func BenchmarkSliceOfNonPointer(b *testing.B) {
	for range b.N {
		DoWithNoPointer()
	}
}

package util

import (
	"fmt"
	"testing"
)

func TestPopEmojis(t *testing.T) {

	fmt.Println(PopEmojis("🌏OKOK🍎"))
}

func BenchmarkPopEmojis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopEmojis("🌏OKOK🍎")
	}
}

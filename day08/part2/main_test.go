package main

import (
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkDoTheThing(b *testing.B) {
	content, _ := os.ReadFile(filepath.Join("..", "input.txt"))
	for i := 0; i < b.N; i++ {
		_ = doTheThing(content)
	}
}

package main

import (
	"testing"
)

func TestBoards(t *testing.T) {
	for i := 0; i < len(boards); i++ {
		b := boards[i]
		width := len(b.colors)
		height := len(b.colors[0])

		for j := 1; j < width; j++ {
			if len(b.colors[j]) != height {
				t.Fatalf("board height - received: %v - expected: %v - index %v", len(b.colors[j]), height, i)
			}
		}

		if len(b.rotation) != width {
			t.Fatalf("rotation width - received: %v - expected: %v - index %v", len(b.rotation), width, i)
		}

		for j := 0; j < width; j++ {
			if len(b.rotation[j]) != height {
				t.Fatalf("rotation height - received: %v - expected: %v - index %v", len(b.rotation[j]), height, i)
			}
		}

	}
}

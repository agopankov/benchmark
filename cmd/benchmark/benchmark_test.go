package main

import (
	"testing"
)

func benchmark(b *testing.B) {
	b.ResetTimer()

	args := []string{"arg1", "arg2", "arg3"}

	for i := 0; i < b.N; i++ {
		PerformTask(args)
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark(b)
}

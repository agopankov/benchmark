package main

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func benchmark(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for k, v := range os.Args[1:] {
			var resultStringSlice []string
			resultStringSlice = append(resultStringSlice, strconv.Itoa(k), v)
			_ = strings.Join(resultStringSlice, ":")
		}
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark(b)
}

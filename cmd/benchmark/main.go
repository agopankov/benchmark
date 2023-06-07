package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func benchmarkVersion1(b *testing.B) {
	args := []string{"arg1", "arg2", "arg3"} // Пример аргументов командной строки

	for i := 0; i < b.N; i++ {
		// Вставьте сюда первую версию кода
		for k, v := range args {
			var resultStringSlice []string
			resultStringSlice = append(resultStringSlice, strconv.Itoa(k), v)
			_ = strings.Join(resultStringSlice, ":")
		}
	}
}

func benchmarkVersion2(b *testing.B) {
	args := []string{"arg1", "arg2", "arg3"} // Пример аргументов командной строки

	for i := 0; i < b.N; i++ {
		// Вставьте сюда вторую версию кода
		for k, v := range args {
			var resultStringSlice []string
			resultStringSlice = append(resultStringSlice, strconv.Itoa(k+1), v)
			_ = strings.Join(resultStringSlice, ":")
		}
	}
}

func main() {
	fmt.Println("Тест производительности для первой версии кода:")
	testing.Benchmark(benchmarkVersion1)

	fmt.Println()

	fmt.Println("Тест производительности для второй версии кода:")
	testing.Benchmark(benchmarkVersion2)
}

package main

import (
	"os"
	"strconv"
	"strings"
)

func PerformTask(args []string) {
	for k, v := range args {
		var resultStringSlice []string
		resultStringSlice = append(resultStringSlice, strconv.Itoa(k), v)
		_ = strings.Join(resultStringSlice, ":")
	}
}

func main() {
	PerformTask(os.Args[1:])
}

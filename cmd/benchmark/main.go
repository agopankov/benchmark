package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	for k, v := range os.Args[1:] {
		var resultStringSlice []string
		resultStringSlice = append(resultStringSlice, strconv.Itoa(k), v)
		_ = strings.Join(resultStringSlice, ":")
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1]
	prefix := GetPrefix(arg, len(arg))
	s := prefix + strings.ToUpper(arg) + prefix

	fmt.Println(s)
}

func GetPrefix(s string, len int) string {
	return strings.Repeat("!", len)
}

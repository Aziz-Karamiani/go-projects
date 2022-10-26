package main

import "fmt"

var i = 10

func main() {
	fmt.Println("Package Level : ", i)
	i := 6
	fmt.Println("Package Level : ", i)
	{
		i := 11
		fmt.Println("Nested Block Level : ", i)
	}
	fmt.Println("Block Level : ", i)
}

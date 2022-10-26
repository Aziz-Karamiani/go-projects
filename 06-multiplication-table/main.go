package main

import "fmt"

func main() {
	var max = 10
	fmt.Printf("%3s |", "X")
	for i := 1; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	fmt.Println("  ------------------------------------------------------")

	for i := 1; i <= max; i++ {
		fmt.Printf("%3d |", i)
		for j := 1; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}

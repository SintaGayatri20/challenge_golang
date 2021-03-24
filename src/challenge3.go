package main

import "fmt"

func main() {
	// IF, ELSE IF, ELSE
	var n int
	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Print((i + 1))
	}
}

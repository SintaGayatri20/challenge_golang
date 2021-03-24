package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

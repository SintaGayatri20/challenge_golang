package main

import "fmt"

func main() {
	// MENGHITUNG LUAS PERMUKAAN TABUNG
	var luas, r, t float32
	const pi = 3.14
	fmt.Print("Masukkan jari-jari tabung: ")
	fmt.Scanln(&r)
	fmt.Print("Masukkan tinggi tabung: ")
	fmt.Scanln(&t)
	luas = 2 * pi * r * (r + t)
	fmt.Println("Luas tabung adalah", luas)
}

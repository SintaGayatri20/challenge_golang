package main

import "fmt"

func main() {
	// IF, ELSE IF, ELSE
	var nilai int
	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&nilai)

	if nilai >= 85 {
		fmt.Println("Grade A")
	} else if nilai >= 75 {
		fmt.Println("Grade B")
	} else if nilai >= 65 {
		fmt.Println("Grade C")
	} else if nilai >= 55 {
		fmt.Println("Grade D")
	} else {
		fmt.Println("TIDAK LULUS !")
	}

}

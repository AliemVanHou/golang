package main

import "fmt"

func faktorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * faktorial(num-1) 
}
}

func main() { 
	var faktorials int
	fmt.Print("Masukkan Angka Faktorial: ")
	fmt.Scanln(&faktorials)
	fmt.Print(faktorial(faktorials))
}
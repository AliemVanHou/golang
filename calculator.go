package main

import (
	"errors"
	"fmt"
)

func calculator(operasi string, bil1, bil2 int) (int, error) {
	result := 0
	switch operasi {
	case "+" :
		result = bil1 + bil2 
	case "-" :
		result = bil1 - bil2
	case "*" :
		result = bil1 * bil2
	case "/" :
		if bil2 == 0 {
			return 0, errors.New("pembagi tidak bisa dengan 0")
		} else {
			result = bil1 / bil2
		}
	}
	return result, nil
}

func main() {
	var operasi string
	var bil1, bil2 int
	fmt.Print("Masukkan Operasi: ")
	fmt.Scanln(&operasi)
		fmt.Printf("Masukkan Angka ke-1: ")
		fmt.Scanln(&bil1)
		fmt.Printf("Masukkan Angka ke-2: ")
		fmt.Scanln(&bil2)

	result, err := calculator(operasi, bil1, bil2)
	if err == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("error", err.Error())
	}

}
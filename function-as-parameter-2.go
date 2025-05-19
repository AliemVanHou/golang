package main

import "fmt"

func check(arr[]int) []string {
	result := []string {}
	for i:=0; i<len(arr); i++ {
		if arr[i] % 2 == 0{
			result = append(result, "genap")
		} else {
			result = append(result, "ganjil")
		}
	}
	return result
}
func checking(arr[]int, filter func([]int)[]string) []string {
	return filter(arr)
}
func main () {
	var input int
	var number int
	container := []int {}
	fmt.Print("Masukkan Banyak Input: ")
	fmt.Scanln(&input)
	for i := 1; i<=input; i++ {
		fmt.Print("Masukkan Angka: ")
		fmt.Scanln(&number)
		container = append(container,number)
	}
	result := checking(container,check)
	for i:=0; i<len(container); i++ {
		fmt.Println(container[i], "adalah bilangan ",result[i] )
	}
}
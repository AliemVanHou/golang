package main

import "fmt"

func main () {
	var num int
	fmt.Print("Put The Number: ")
	fmt.Scanln(&num)
	fmt.Println(num, "Adalah Bilangan", checking1(num, checking) )
}

func checking(num int) string {
	if num % 2  == 0 {
		return "genap"
	} else {
		return "ganjil"
	}
}

type filter func(int) string
func checking1(num int, filter filter) string{
	return filter(num)
}
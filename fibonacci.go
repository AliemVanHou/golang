package main

import "fmt"

func fibonacci(num int) []int {
	a := 0
	b := 1
	c := 0
	result := []int {}
	for i:=0; i<num; i++ {
		result = append(result, a)
		c = a + b
		a = b
		b = c
	}
	return result
}
func main() {
	input := 0

	fmt.Print("Masukkan Deret Fibonacci: ")
	fmt.Scanln(&input)

	fmt.Print(fibonacci(input))
}
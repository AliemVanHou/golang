package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address := Address {
		City: "South Tangerang",
		Province: "Banten",
		Country: "Indonesia",
	}
	address2 := &address
	address2.City = "Tangerang"

	fmt.Println(address)
	fmt.Println(address2)
	
}
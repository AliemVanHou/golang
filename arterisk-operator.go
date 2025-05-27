package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address := Address {
		City: "Tangerang",
		Province: "Banten",
		Country: "Indonesia",
	}
	address2 := &address
	address2.City = "South Tangerang"
	
	fmt.Println(address)
	fmt.Println(address2)

	*address2 = Address {
		City: "Jakarta",
		Province: "Jakarta",
		Country: "Indonesia",
	}
	fmt.Println(address)
	fmt.Println(address2)
}
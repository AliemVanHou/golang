package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{
		City:     "South Tangerang",
		Province: "Banten",
		Country:  "Korea",
	}
	fmt.Println(address)
	(changeCountryToIndonesia(&address))
	fmt.Println(address)
}
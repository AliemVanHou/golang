package main

import "fmt"

func random() any {
	return 12
}

func main() {
	result := random()
	
	switch value := result.(type) {
	case string :
		fmt.Println("String", value)
	case int :
		fmt.Println("Integer", value)
	default :
		fmt.Println("Unknown") 
	}
}
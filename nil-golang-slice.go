package main

import "fmt"

func mySlice(name []string) []string {
	if len(name) == 0 {
		return nil
	} else {
		return name
	}
}

func main() {
	hallo := []string {}
	if hallo == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(mySlice(hallo))
	}
}
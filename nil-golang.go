package main

import "fmt"

func myMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"Name ": name,
		}
	}
}

func main() {
	data := myMap("Reinaldi Alimsyah")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}
}
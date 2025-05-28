package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation error") 
	ErrNotFound = errors.New("not found error")
)

func getById(id string) error {
if id == "" {
		return ErrValidation
	} else if id != "Reinaldi Alimsyah" {
		return ErrNotFound 
	}
	return nil
}

func main() {
	err := getById("")
	if err != nil {
		if errors.Is(err, ErrValidation) {
			fmt.Println("validation errors")
		} else if errors.Is(err, ErrNotFound) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	} else {
		fmt.Println("Sukses")
	}
	}
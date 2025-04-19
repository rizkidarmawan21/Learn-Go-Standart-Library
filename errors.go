package main

import "errors"

var (
	ValidationError = errors.New("Validation error")
	NotFoundError   = errors.New("Not found error")
)

func getById(id int) error {
	if id <= 0 {
		return ValidationError
	}
	if id > 100 {
		return NotFoundError
	}
	return nil
}

func main() {
	err := getById(0)
	if err != nil {
		if errors.Is(err, ValidationError) {
			println("Validation error occurred")
		} else if errors.Is(err, NotFoundError) {
			println("Not found error occurred")
		}
	} else {
		println("No error occurred")
	}
}

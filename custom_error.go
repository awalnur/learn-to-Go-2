package main

import "fmt"

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(message string) *CustomError {
	return &CustomError{message}
}

func main() {
	if 10 > 3 {

		error := NewCustomError("Error Message")
		fmt.Println(error)
	}
}

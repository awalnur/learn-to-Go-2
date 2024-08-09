package main

import (
	"errors"
	"fmt"
)

func Pembagian(a int, b int) (int, error) {
	if a > b {
		return 0, errors.New("this Error")
	} else {
		return a, nil
	}
}
func main() {

	nilai, error := Pembagian(12, 11)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(nilai)
	}

}

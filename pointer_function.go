package main

import "fmt"

type Address struct {
	City, State string
	Width       int
}

func ChangeCountry(addr *Address) {

	// * menandakan asteris
	addr.City = "Jakarta"
}
func main() {
	adr := Address{City: "Bandung", State: "cjt", Width: 1000}
	fmt.Println("Address is", adr)
	adr2 := adr   // Copy
	adr3 := &adr2 // Reference

	ChangeCountry(&adr)

	adr3.State = "Bebas"
	fmt.Println(adr)
	fmt.Println(adr2)
	fmt.Println(adr3)

	//adr3 = &Address{City: "Semarang", State: "cjt", Width: 1000}
	//fmt.Println(adr3)

	*adr3 = Address{City: "Bandung", State: "cjt", Width: 1000} // asteris
	fmt.Println(adr2)
}

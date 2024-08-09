package main

import "fmt"

type HashName interface {
	GetName() string
}

func SayHellow(str HashName) {
	fmt.Println("hello", str.GetName())
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

func main() {
	person := Person{Name: "Person Name"}
	SayHellow(person)

}

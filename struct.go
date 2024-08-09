package main

import "fmt"

type Customer struct {
	Name string
	Age  int
}

func (cust Customer) sayHello(name string) string {
	//fmt.Println(cust.Name, name)
	return "Hello " + cust.Name
}

func main() {
	cs := Customer{Name: "test"}
	fmt.Println(cs.sayHello("tess"))
}

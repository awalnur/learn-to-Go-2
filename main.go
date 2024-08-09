package main

import (
	database "belajar-golang/database"
	hels "belajar-golang/helper"
	_ "belajar-golang/other"
	"fmt"
)

func main() {
	hel := hels.SayHello("Hello World")
	fmt.Println(hels.AppName, hels.Public())
	fmt.Println(hel)
	// Init function
	con := database.GetConnection()
	fmt.Println(con)

	// blank identifier -> ketika import package tanpa mau menggunakannya

}

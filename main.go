package main

import (
	"fmt"
	"tudai21/services"
)

func main() {
	var s services.StringFormat
	/* c := "TX08ABC" */
	var c string
	fmt.Println("Ingrese la cadena")
	fmt.Scanf("%s\n", &c)
	sf, err := s.ParseString(c)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
	}
	fmt.Println(sf)
}

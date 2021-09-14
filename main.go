package main

import (
	"fmt"
	"tudai21/services"
)

func main() {
	var s services.Result
	/* c := "TX08ABC" */
	var c string
	fmt.Println("Ingrese la cadena")
	fmt.Scanf("%s\n", &c)
	r, err := s.ParseString(c)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Print(err)
	} else {
		fmt.Println(r)
	}

}

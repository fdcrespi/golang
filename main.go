package main

import (
	"fmt"
	"tudai21/services"
)

func main() {
	var s services.Result
	var c string
	//c := "TX08ABC"
	fmt.Println("Ingrese la cadena")
	fmt.Scanf("%s\n", &c)
	r, err := s.GenerateString(c)
	if r == nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

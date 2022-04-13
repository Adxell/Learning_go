package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Algo no se cumple ")
	}

	// With or

	if valor1 == 0 || valor2 == 2 {
		fmt.Println("algo si se cumple")
	} else {
		fmt.Println("ninguna se cumple")
	}

	// convertir texto a numero 

	value, err := strconv.Atoi("53")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value: ", value)
}

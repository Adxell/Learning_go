package main

import "fmt"

func main() {
	//Declaracion de variables
	helloMessage := "hello"
	worldMessage := "world"

	// Println para colocar un salto de linea
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf

	nombre := "Adxell"
	age := 21

	fmt.Printf("%s tiene la edad de %d years\n", nombre, age)
	fmt.Printf("%v tiene la edad de %v years\n", nombre, age)
	//%s para string
	//%d para tipo numerico
	//%v para cuando no sabemos que tipo de datos va a tener las variables

	//Sprintf

	message := fmt.Sprintf("%s tiene la edad de %d years", nombre, age)
	fmt.Println(message)

	//Tipo datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("Edad: %T\n", age)
	//Para ver el tipo de dato que es

}

package main

import (
	"fmt"
)

func main() {

	//Area de un rectangulo

	base := 10
	altura := 12
	areaRetangulo := base * altura
	fmt.Println("Area del rectangulo:", areaRetangulo)

	//Area trapecio
	var base1 float64 = 6
	var altura1 float64 = 4
	var base2 float64 = 3

	areaTrapecio := altura1 * ((base1 + base2) / 2)

	fmt.Println("Area del trapecio:", areaTrapecio)

	//Area circulo

	const pi float64 = 3.1416
	var radio float64 = 5

	areaCirculo := pi * (radio * radio)
	fmt.Println("Area del circulo:", areaCirculo)

}

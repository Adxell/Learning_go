package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	//Declaracion de variables enteras

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero por default
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	//operadores
	x := 10
	y := 50

	// Suma
	resultSuma := x + y
	fmt.Println("Suma:", resultSuma)

	// Resta
	resultResta := y - x
	fmt.Println("Resta:", resultResta)

	// Multiplicacion
	resultMultiplicacion := x * y
	fmt.Println("Multiplicacion:", resultMultiplicacion)

	//Division
	resultDivision := y / x
	fmt.Println("Division:", resultDivision)

	//Modulo
	resultModulo := y % x
	fmt.Println("Modulo:", resultModulo)

	// Incrementar 
	x++
	fmt.Println("Imcremental:", x)

	//Decrementar
	x--
	fmt.Println("Decrementar:", x)
}

package main

import (
	"fmt"
	"math"
)

func calcularAreaRectagulo(base, altura float64) float64 {
	return base * altura
}

func calcularAreaTrapecio(base1, base2, altura float64) float64 {
	return altura * ((base1 + base2) / 2)
}

func calcularAreaCirculo(r float64) float64 {
	return (r * r) * math.Pi
}

func main() {
	//Area Rectagulo
	areaRectangulo := calcularAreaRectagulo(4.5, 2)
	fmt.Println("Area rectangulo:", areaRectangulo)

	//Area trapecio
	areaTrapecio := calcularAreaTrapecio(6, 3, 4)
	fmt.Println("Area trapecio:", areaTrapecio)

	//Area circulo
	areaCirculo := calcularAreaCirculo(5)
	fmt.Println("Area circulo:", areaCirculo)
}

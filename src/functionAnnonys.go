package main

import "fmt"

//Declarar funcion 
func normalFunction(message string){
	fmt.Println(message)
}
func tripleArgument(a, b int, c string){
	fmt.Println(a, b, c)
}

func returnValue( a int) int{
	return a * 2
}
func doblueReturn(a int) (c , d int){
	return a, a*2
}

func main() {
	normalFunction("Hola Adxell")
	tripleArgument(1,2,"Hola como estas?")

	value := returnValue(2)
	fmt.Println("value:", value)

	value1, _ := doblueReturn(2)
	//fmt.Println("value 1 y value 2:", value1, value2)
	fmt.Println("value 1:", value1)
}

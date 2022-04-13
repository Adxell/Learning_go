package main

import (
	"fmt"
)

func pairNumber(n int) bool {
	return n%2 == 0
}
func userSystem(user, password string) bool {
	return user == "Adxell" && password == "Adxell1234567"
}

func main() {
	if pairNumber(11) {
		fmt.Println("your number is pair")
	} else {
		fmt.Println("your number is not pair")
	}
	if userSystem("Adxell", "Adxell123456") {
		fmt.Println("Access successful")
	} else {
		fmt.Println("Access Denied")
	}
}

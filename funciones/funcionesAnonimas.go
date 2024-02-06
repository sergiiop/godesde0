package funciones

import "fmt"

func Calculos() {
	suma := func(number1 int, number2 int) int {
		return number1 + number2
	}

	fmt.Println(suma(10, 25))
}
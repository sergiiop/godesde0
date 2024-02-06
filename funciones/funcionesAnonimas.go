package funciones

import "fmt"

func Calculos() {

	var numero3 int = 21
	var numero4 int = 221

	calculo := func(number1 int, number2 int) int {
		return number1 + number2 + numero3
	}
	fmt.Println(calculo(10, 25))


	calculo = func(number1 int, number2 int) int {
		return number1 + number2 + numero3 + numero4
	}

	fmt.Println(calculo(10, 25))
}
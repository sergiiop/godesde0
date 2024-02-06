package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplicationTable() string {

	scanner := bufio.NewScanner(os.Stdin)

	var firstNumber int
	var err error

	var texto string

	for {
		fmt.Println("Ingrese número: ")
		if scanner.Scan() {
			firstNumber, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Ingresa un numero valido")
				continue
			}
			break
		}
	}

	for i := 0; i < 10 ; i ++ {
		texto += fmt.Sprintf("%d * %d = %d\n", firstNumber, i+1, firstNumber * (i+1))
	}

	return texto
}
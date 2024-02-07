package main

import (
	"github.com/sergiiop/godesde0/arreglos_slices"
)

// "runtime"
// "github.com/sergiiop/godesde0/variables"

// runtime obtiene toda la información del sistema

func main() {
	// estado, texto := variables.ConviertoATexto(144)

	// fmt.Println(estado)
	// fmt.Println(texto)

	// // Asignando y haciendo la pregunta, separando por ;

	// // Operadores de comparación

	// // ==; >; <; >= ; <= ; && ; ||
	// if os := runtime.GOOS; os =="linux" || os =="OS X." {
	// 	fmt.Println("Esto no es windows, es ", os)
	// } else {
	// 	fmt.Println("Esto es Window")
	// }

	// switch os:= runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	// formattedNumber, message, err := ejercicios.FormattedString("12w3")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message, formattedNumber)

	// teclado.IngresoNumeros()

	// for {
	// 	fmt.Println("Hola")
	// 	// break <- romper el for
	// }
	// fmt.Println(ejercicios.MultiplicationTable())
	// funciones.Exponencia(2)
	// arreglos_slices.MuestroArreglos()
	arreglos_slices.Capacidad()
}
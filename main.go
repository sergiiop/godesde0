package main

import (
	"fmt"

	"github.com/sergiiop/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(144)

	fmt.Println(estado)
	fmt.Println(texto)
}
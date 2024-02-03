package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string // Esta sera exportada por estar fuera de la función por tener mayúscula y sera visto por cualquier archivo del paquete y por los que importen el paquete
var apellido = "Perez" // Esta no sera exportada aunque este fuera de una function ya que esta en minúscula
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Sergio"
	Estado = true
	Sueldo = 153.5
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}


// Si la function devuelve mas de un dato se debe establecer entre paréntesis y separados por , (,) eje: (bool, string)
func ConviertoATexto(numero int) (bool, string) {

	texto := strconv.Itoa(numero)

	return true, texto
}
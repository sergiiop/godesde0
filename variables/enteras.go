package variables

import "fmt"

// Para que una variable o una función pueda ser accedida fuera del paquete esta debe iniciar con mayúsculas
func MuestroEnteros() {
	var intcomun int // Establecer el tipo de dato, cuando se establece sin asignación, no se le asigna null, se asigna el valor mínimo de ese tipo de dato en este caso sera 0

	intde32 := int32(10) // El tipo de dato se establece de la asignación en este caso int32

	intde64 := int64(100) // tipo de dato int64

	// intde64 = intcomun(100) // no se puede asignar otro tipo de variable a una que ya se le estableció el tipo

	fmt.Println("int común = ", intcomun)
	fmt.Println("int de 32", intde32)
	fmt.Println("int de 64", intde64)
}
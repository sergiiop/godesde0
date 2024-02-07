package arreglos_slices

import (
	"fmt"
)

// Una de las diferencias entre slice y un vector, es que los slices no tienen una cantidad de dimension de elementos
//slice
// var tabla []int

// vector
// var tabla [10]int

// var tabla[10]int
var tabla[10]int = [10]int{10,0,52}

func MuestroArreglos () {
	tabla[8] = 2
	tabla[1] = 1

	tabla2 := [10]string{"pablo", "juanchin"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
}
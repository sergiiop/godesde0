package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 4, 5}

var myArreglo [10]int = [10]int{10,21,21,4,21,4,5, 6, 9}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := myArreglo[3:] // Slice creado desde un vector desde la posicion 3 hasta el final

	porcion2 := myArreglo[:5] // Slice creado desde un vector desde la posicion 0 hasta la 5
	porcion3 := myArreglo[5:7] // Slice creado desde un vector desde la posicion 5 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20) // size 5, capacidad 20
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos)) // Con cap vemos la capacidad de un slice y con len su size

	nums := make([]int, 0, 0)

	for i := 0; i< 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums)) // Con cap vemos la capacidad de un slice y con len su size

}
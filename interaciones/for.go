package interaciones

import (
	"fmt"
)

func Iterar() {
	// i := 0

	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	for i := 0; i<10 ; i+=1 {

		if i == 6 {
			break // Rompe el for
		}

		if i == 2 {
			continue // Cuando i sea 2 va a saltar a la siguiente iteración
		}

		fmt.Println(i)
	}
}
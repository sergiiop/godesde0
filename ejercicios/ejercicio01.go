package ejercicios

import (
	"errors"
	"strconv"
)

func FormattedString(valor string) (int, string, error) {
	formattedEntero, errFormat := strconv.Atoi(valor)

	if errFormat != nil {
		return 0,"",errors.New("El valor suministrado no es valido")
	}

	if formattedEntero > 100 {
		return formattedEntero, "Es mayor a 100", nil
	} else {
		return formattedEntero, "Es menor a 100", nil
	}
}
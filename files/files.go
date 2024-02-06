package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sergiiop/godesde0/ejercicios"
)

var fileName string = "./files/archivos/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.MultiplicationTable()


	// Esta función crea un archivo, si el archivo existe borra el anterior y lo crea nuevamente
	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo " +err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	// todo archivo que se abre para grabar o lectura se tiene que cerrar
	archivo.Close()


}

func SumaTabla() {
	var texto string = ejercicios.MultiplicationTable()

	if !Append(fileName, texto){
		fmt.Println("Error al concatenar el contenido")
	}
}

func Append(filen string, texto string) bool {
	// os.Open() abre el archivo solo de lectura
	// con os.O_WRONLY le decimos que habra el archivo en modo escritura pero va a borrar el contenido que tenga, pero con os.O_APPEND le decimos que no borre el contenido si no que adicione al final
	arch, err := os.OpenFile(filen, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el append " +err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el write " +err.Error())
		return false
	}

	arch.Close()

	return true
}

func LeoArchivo() {
	arch, err := os.Open(fileName)

	if err != nil{
		fmt.Println("Hubo un error " + err.Error())
		return
	}

	scanner := bufio.NewScanner(arch)

	for scanner.Scan() {
		registro := scanner.Text()

		fmt.Println("> "+ registro)
	}

	arch.Close()
}
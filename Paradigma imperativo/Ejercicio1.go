package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	/*1)Haga un programa que cuente e indique el número de caracteres, el número de palabras y el número de líneas
	de un texto cualquiera (obtenido de cualquier forma que considere). Considere hacer el cálculo de las tres
	variables en el mismo proceso.*/

	var caracteres int = 0
	var palabras int = 0
	var lineas int = 1
	var texto string = "Él me mintió \nÉl me dijo que me amaba y no era verdad \nÉl me mintió \nNo me amaba nunca me amó"

	// Recorrer el texto caracter por caracter
	for _, c := range texto {
		caracteres++
		if c == ' ' {
			palabras++
		}
		if c == '\n' {
			lineas++
		}
	}
	// Suma la última palabra que no detecta porque deja de contar en el último espacio
	palabras++
	fmt.Println("Texto de prueba:")
	fmt.Println(texto)
	fmt.Println("\n"+"Numero de caracteres:", caracteres)
	fmt.Println("Numero de palabras:", palabras)
	fmt.Println("Numero de lineas:", lineas)
}

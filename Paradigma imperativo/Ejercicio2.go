package main

import (
	_ "fmt"
)

func main() {
	/*2)Escriba el programa más eficiente que pueda para imprimir en pantalla la siguiente figura, Para dicho fin,
	  escriba y use una función que reciba de parámetro la cantidad de elementos de la línea del centro, la cual debe
	  ser impar positiva.*/
	imprimirFigura(5)
}

func imprimirFigura(n int) {
	if n%2 == 0 || n < 0 {
		println("La cantidad de elementos de la línea del centro debe ser un número impar y positivo")
	} else {
		for i := 0; i < n; i++ {
			var numAsteriscos int
			if i <= n/2 {
				numAsteriscos = 2*i + 1
			} else {
				numAsteriscos = 2*(n-i) - 1
			}
			numEspacios := (n - numAsteriscos) / 2
			println(espacios(" ", numEspacios), asteriscos("*", numAsteriscos))
		}
	}
}

func espacios(str string, n int) string {
	var figura string
	for i := 0; i < n; i++ {
		figura += str
	}
	return figura
}

func asteriscos(str string, n int) string {
	var figura string
	for i := 0; i < n; i++ {
		figura += str
	}
	return figura
}

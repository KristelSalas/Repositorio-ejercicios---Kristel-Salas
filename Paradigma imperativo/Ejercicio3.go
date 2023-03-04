package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	/*3)Escriba una función que permita rotar una secuencia de elementos N posiciones a la izquierda o a la
	derecha según sea el caso en función al parámetro que se reciba. Los parámetros deben ser un puntero a un
	arreglo previamente creado, la cantidad de movimiento de rotación y la dirección (0 = izq y 1 = der).*/
	arreglo := []int{1, 2, 3, 4, 5, 6, 7, 8}
	pun := &arreglo
	fmt.Println("Arreglo inicial: ", arreglo)
	rotar(pun, 5, 1)
}

func rotar(arr *[]int, mover int, direccion int) {
	length := len(*arr)
	if direccion == 0 {
		for i := 0; i < mover; i++ {
			temporal := (*arr)[0]
			copy((*arr)[0:length-1], (*arr)[1:])
			(*arr)[length-1] = temporal
		}
	} else {
		for i := 0; i < mover; i++ {
			temporal := (*arr)[length-1]
			copy((*arr)[1:], (*arr)[0:length-1])
			(*arr)[0] = temporal
		}
	}
	fmt.Println("Arreglo modificado: ", *arr)
}

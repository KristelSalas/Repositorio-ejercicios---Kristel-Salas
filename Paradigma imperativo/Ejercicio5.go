package main

import (
	"fmt"
	"sort"
)

type Producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []Producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
// de elementos del producto y eventualmente el precio si es que es diferente
func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 {
		(*l)[prod].cantidad = (*l)[prod].cantidad + cantidad
	} else {
		*l = append(*l, Producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}

}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

// modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar
// tambein que reciba el numero de articulos a vender
func (l *listaProductos) venderProducto(nombre string, cantidad int) {
	var prod = l.buscarProducto(nombre)
	if (*l)[prod].cantidad < cantidad {
		*l = append((*l)[:prod], (*l)[prod+1:]...)
	} else {
		(*l)[prod].cantidad = (*l)[prod].cantidad - cantidad
	}
}
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("arroz", 5, 2500)
	lProductos.agregarProducto("chocolate", 3, 1000)
	lProductos.agregarProducto("chocolate", 3, 1000)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)

	lProductos.venderProducto("arroz", 19)
}

// debe retornar una nueva lista con productos con existencia mínima
func (l *listaProductos) listarProductosMinimos() listaProductos {
	var productosMinimo listaProductos
	for _, producto := range *l {
		if producto.cantidad < existenciaMinima {
			productosMinimo = append(productosMinimo, producto)
		}
	}
	return productosMinimo
}

/*
A partir de la lista de productos con mínimas existencias de inventario, ampliar el inventario con la compra
de más unidades de cada producto de esta lista hasta que cumplan con el mínimo establecido.
*/
func (l *listaProductos) aumentarInventarioDeMinimos() {
	listaMinimos := lProductos.listarProductosMinimos()
	for _, prod := range listaMinimos {
		cant := existenciaMinima - prod.cantidad
		prod.cantidad = cant
		l.agregarProducto(prod.nombre, prod.cantidad, prod.precio)
	}
}

/*
Crear una función que ordene la lista de productos usando como llave para el ordenamiento cualquiera de
los elementos de la estructura producto. La lista/slice debe quedar modificada al finalizar el método.
Se solicita investigar y hacer uso de alguna función predefinida de alguna librería del lenguaje Go que
ayude a resolver el problema
*/
func (l *listaProductos) ordenarProductos() listaProductos {
	sort.Slice(lProductos, func(i, j int) bool {
		return lProductos[i].precio < lProductos[j].precio
	})
	return lProductos
}

func main() {
	llenarDatos()
	fmt.Println("Lista inicial de productos:", "\n", lProductos)
	fmt.Println("\n", "Productos con existencia por debajo del minimo:", "\n", lProductos.listarProductosMinimos())
	lProductos.aumentarInventarioDeMinimos()
	fmt.Println("\n", "Lista de productos con el minimo aumentado: ", "\n", lProductos)
	lProductos.ordenarProductos()
	fmt.Println("\n", "Productos ordenados por precio de manera ascendente: ", "\n", lProductos)
}

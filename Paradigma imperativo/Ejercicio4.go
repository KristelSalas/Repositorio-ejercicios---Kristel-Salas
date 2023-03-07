package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type producto struct {
	nombre      string
	descripcion string
	montoVenta  int32
}

type Productos []producto

var factura Productos

const rangoPagoImpuestos = 20000
const porcentajeImpuesto = 0.13

func (f *Productos) agregarProducto(nom string, desc string, pre int32) {
	idx := slices.IndexFunc(*f, func(p producto) bool { return p.nombre == nom })
	if idx == -1 {
		*f = append(*f, producto{nom, desc, pre})
	} else {
		fmt.Println("cant add existing product to the list")
	}
}

func (f *Productos) calcularImpuestoFactura() int32 {
	lista := filter1(*f, func(p producto) bool {
		return p.montoVenta > rangoPagoImpuestos
	})
	lista2 := map2(lista, func(p producto) int32 {
		return int32(float64(p.montoVenta) * porcentajeImpuesto)
	})
	return reduce(lista2)
}

func (f *Productos) calcularMontoFactura() int32 {
	lista := map2(*f, func(p producto) int32 {
		return p.montoVenta
	})
	return reduce(lista) //reduce hace una suma total
}
func (f *Productos) calcularMontoFactura2() int32 {
	lista := map1(*f, func(p producto) int32 {
		return p.montoVenta
	})
	return reduce(lista)
}

// ---------------------------------------------------------------------------------------------------
// funciones map y filter para aplicaciones específicas

func map1(list Productos, f func(producto) int32) []int32 {
	mapped := make([]int32, len(list))

	for i, e := range list {
		mapped[i] = f(e)
	}
	return mapped
}

func filter1(list Productos, f func(producto) bool) Productos {
	filtered := make(Productos, 0)

	for _, element := range list {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

// ---------------------------------------------------------------------------------------------------
// Funciones map y filter genericas PRACTICA

// Recibe
func map2[p1, p2 any](list []p1, f func(p1) p2) []p2 {
	mapped := make([]p2, len(list))
	for i, e := range list {
		mapped[i] = f(e)
	}
	return mapped
}

func filter2[p1 any](list []p1, f func(p1) bool) []p1 {
	filtered := make([]p1, 0)

	for _, element := range list {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

//---------------------------------------------------------------------------------------------------

func reduce(list []int32) int32 {
	var result int32 = 0
	for _, v := range list {
		result += v
	}
	return result
}

// Construir la versión con tipos genéricos de las tres funciones anteriores map/filter para que acepten slices y funciones de cualquier otro tipo
// AYUDA SOBRE TIPOS GENERICOS: https://gobyexample.com/generics

// // Probar su funcionamiento creando una lista/slice de personas y filtrando aquellas personas que sean mayores de edad
// EJERCICIO PARA REPO

type persona struct {
	nombre   string
	apellido string
	edad     int
}

type Personas []persona

var listaPersonas Personas

func (f *Personas) agregarPersona(nom string, ape string, edad int) {
	*f = append(*f, persona{nom, ape, edad})
}

func main() {
	/*factura.agregarProducto("tarjeta madre", "Asus", 54200)
	factura.agregarProducto("mouse", "alámbrico", 15000)
	factura.agregarProducto("teclado", "gamer con luces", 30000)
	factura.agregarProducto("memoria ssd", "2 gb", 65000)
	factura.agregarProducto("cable video", "display port 1m", 18000)

	println(factura.calcularMontoFactura())
	println(factura.calcularImpuestoFactura())*/

	listaPersonas.agregarPersona("Mario", "Salas", 20)
	listaPersonas.agregarPersona("Maria", "Gomez", 17)
	listaPersonas.agregarPersona("Lucia", "Monteverde", 21)
	listaPersonas.agregarPersona("Alba", "Villegas", 15)
	listaPersonas.agregarPersona("Pepe", "Arguedas", 24)
	listaPersonas.agregarPersona("Jose", "Acuña", 14)

	fmt.Println("Lista original:", listaPersonas)
	fmt.Println("Lista de Mayores", filter2(listaPersonas, func(p persona) bool { return p.edad > 18 }))
}

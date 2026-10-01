package main

import "fmt"

var nombres []string
var subtotales []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)
	nombres = append(nombres, nombre)
	subtotales = append(subtotales, subtotal)
	fmt.Println("Venta registrada. Subtotal:", subtotal)
}

func MostrarEstadisticas() {
}

func main() {
	var opcion int

	for opcion != 3 {
		fmt.Println("1. Registrar venta")
		fmt.Println("2. Mostrar estadisticas")
		fmt.Println("3. Salir")
		fmt.Scan(&opcion)

		if opcion == 1 {
			fmt.Println("1. Arroz - 1.25")
			fmt.Println("2. Leche - 0.95")
			fmt.Println("3. Pan - 0.50")

			var producto, cantidad int
			fmt.Scan(&producto)
			fmt.Scan(&cantidad)

			if producto == 1 {
				RegistrarVenta("Arroz", 1.25, cantidad)
			} else if producto == 2 {
				RegistrarVenta("Leche", 0.95, cantidad)
			} else if producto == 3 {
				RegistrarVenta("Pan", 0.50, cantidad)
			} else {
				fmt.Println("Producto no valido")
			}
		} else if opcion == 2 {
			MostrarEstadisticas()
		}
	}
}
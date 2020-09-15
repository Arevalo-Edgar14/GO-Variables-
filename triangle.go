package main

import "fmt"

func main() {
	var base float32   // float to return a float
	var height float32 // float to return a float
	fmt.Println("Ingresa cuanto mide la base")
	fmt.Scan(&base) // Scan to read only an input if not height read isn't make
	fmt.Println("Ingresa cuanto mide la altura")
	fmt.Scan(&height)
	area := base * height / 2 // creation of new variable with the value of the operation will recibe the data type of the asignation
	fmt.Println("El area del triangulo es: ", area)
}

package main

import "fmt"

func main() {
	fmt.Println("Ingresa cuanto mide uno de los lados del cuadrado")
	var side int
	fmt.Scanf("%d", &side)
	fmt.Println("El area del cuadrado es: ", side*side)
}

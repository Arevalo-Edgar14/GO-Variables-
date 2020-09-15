package main

import (
	"fmt"
	"math"
)

func main() {
	var radio float64
	fmt.Println("Ingresa cuanto mide el radio")
	fmt.Scanf("%f", &radio)
	area := math.Pi * radio * radio
	fmt.Println("El area del circulo es: ", area)
}

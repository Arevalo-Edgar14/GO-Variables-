package main

import "fmt"

func main() {
	var farenheit float64
	fmt.Println("Ingresa es la temperatura farenheit")
	fmt.Scanf("%f", &farenheit)
	const farenheitCeroGradesCelcius = 32.
	const factor = 5. / 9. //for constants get the datatype of the types on the operationthat's why 5. (to work with float as is a fraction)
	celcius := (farenheit - farenheitCeroGradesCelcius) * factor
	fmt.Println("La temperatura en grados celcius es: ", celcius)
}

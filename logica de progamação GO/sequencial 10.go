package main

import (
	"fmt"
)

func main() {
	var pesoKg float64

	// Solicitar o peso em quilogramas
	fmt.Print("Digite o peso da pessoa em quilogramas: ")
	fmt.Scan(&pesoKg)

	// Converter para libras
	pesoLbs := pesoKg / 0.45359237

	fmt.Printf("O peso da pessoa em libras é: %.2f lbs\n", pesoLbs)
}

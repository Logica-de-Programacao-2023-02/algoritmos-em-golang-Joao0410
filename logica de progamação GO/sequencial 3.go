package main

import "fmt"

func main() {
	var peso, altura, imc float64

	// Solicitando o peso ao usuário
	fmt.Print("Digite o seu peso (kg):")
	fmt.Scan(&peso)

	// Solicitando a altura ao usuário
	fmt.Print("Digite a sua altura (m): ")
	fmt.Scan(&altura)

	// Calculando o IMC
	imc = peso / (altura * altura)

	// Mostrando o IMC
	fmt.Printf("Seu IMC é: %.2f\n", imc)
}

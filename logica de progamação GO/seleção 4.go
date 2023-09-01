package main

import "fmt"

func main() {
	var altura float64
	var sexo string
	var pesoIdeal, pesoAtual float64

	// Solicitar altura, sexo e peso atual
	fmt.Println("Digite sua altura (em metros, por exemplo, 1.75): ")
	fmt.Scan(&altura)

	fmt.Println("Digite seu sexo (M para masculino, F para feminino): ")
	fmt.Scan(&sexo)

	fmt.Println("Digite seu peso atual (em kg): ")
	fmt.Scan(&pesoAtual)

	// Calcular o peso ideal
	if sexo == "M" || sexo == "m" {
		pesoIdeal = (72.7 * altura) - 58
	} else if sexo == "F" || sexo == "f" {
		pesoIdeal = (62.1 * altura) - 44.7
	} else {
		fmt.Println("Sexo inválido.")
		return
	}

	// Determinar se a pessoa está abaixo, dentro ou acima do peso ideal
	if pesoAtual < pesoIdeal {
		fmt.Println("Você está abaixo do peso ideal.")
	} else if pesoAtual == pesoIdeal {
		fmt.Println("Você está no peso ideal.")
	} else {
		fmt.Println("Você está acima do peso ideal.")
	}
}

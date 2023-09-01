package main

import "fmt"

func main() {
	var diasTrabalhados int
	var valorDiaria float64

	// Solicitar o número de dias trabalhados
	fmt.Print("Digite o número de dias trabalhados pelo funcionário: ")
	fmt.Scan(&diasTrabalhados)

	// Solicitar o valor da diária
	fmt.Print("Digite o valor da diária: R$")
	fmt.Scan(&valorDiaria)

	// Calcular o salário
	salario := float64(diasTrabalhados) * valorDiaria

	fmt.Printf("O salário do funcionário é: R$%.2f\n", salario)
}

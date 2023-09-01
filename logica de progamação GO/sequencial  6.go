package main

import "fmt"

func main() {
	var salario float64

	// Solicitar o salário atual do funcionário
	fmt.Print("Digite o salário atual do funcionário: R$")
	fmt.Scan(&salario)

	// Calcular o aumento
	aumento := salario * 0.15
	novoSalario := salario + aumento

	fmt.Printf("Com um aumento de 15%%, o novo salário é: R$%.2f\n", novoSalario)
}

package main

import "fmt"

func main() {

	var salario, novoSalario float64

	// Solicitar o salário atual do funcionário
	fmt.Print("Digite o salário atual do funcionário: R$")
	fmt.Scan(&salario)

	// Verificar o salário e calcular o aumento apropriado
	if salario < 1000 {
		novoSalario = salario * 1.10
	} else {
		novoSalario = salario * 1.05
	}

	fmt.Printf("O novo salário é: R$%.2f\n", novoSalario)
}

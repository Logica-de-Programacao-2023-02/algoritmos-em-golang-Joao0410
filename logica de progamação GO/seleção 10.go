package main

import (
	"fmt"
)

func main() {
	var idade int

	// Solicitar a idade da pessoa
	fmt.Print("Digite a idade da pessoa: ")
	fmt.Scan(&idade)

	// Classificar a idade de acordo com a tabela
	switch {
	case idade <= 9:
		fmt.Println("Classificação: mirim")
	case idade >= 10 && idade <= 13:
		fmt.Println("Classificação: infantil")
	case idade >= 14 && idade <= 17:
		fmt.Println("Classificação: juvenil")
	case idade >= 18:
		fmt.Println("Classificação: adulto")
	default:
		fmt.Println("Idade inválida!")
	}
}

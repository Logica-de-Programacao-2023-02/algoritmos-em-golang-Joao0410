package main

import "fmt"

func main() {
	var preco float64

	// Solicitar o preço do produto
	fmt.Print("Digite o preço do produto: R$")
	fmt.Scan(&preco)

	// Calcular o desconto
	desconto := preco * 0.10
	precoComDesconto := preco - desconto

	fmt.Printf("O valor do produto com desconto de 10%% é: R$%.2f\n", precoComDesconto)
}

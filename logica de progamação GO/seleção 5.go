package main

import (
	"fmt"
)

func main() {
	var numero int

	// Solicitar o número
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	// Verificar se o número é múltiplo de 3 e de 5 ao mesmo tempo
	if numero%15 == 0 {
		fmt.Println("O número é múltiplo de 3 e de 5 ao mesmo tempo.")
	} else {
		fmt.Println("O número não é múltiplo de 3 e de 5 ao mesmo tempo.")
	}
}

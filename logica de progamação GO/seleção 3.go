package main

import "fmt"

func main() {
	var numero int

	//  número
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	// par ou ímpar
	if numero%2 == 0 {
		fmt.Println("O número é par.")
	} else {
		fmt.Println("O número é ímpar.")
	}
}

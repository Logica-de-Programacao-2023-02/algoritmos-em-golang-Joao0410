package main

import "fmt"

func main() {
	var numero int

	// Solicitar o número
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	// Calcular o antecessor e o sucessor
	antecessor := numero - 1
	sucessor := numero + 1

	fmt.Printf("Antecessor: %d\n", antecessor)
	fmt.Printf("Sucessor: %d\n", sucessor)
}

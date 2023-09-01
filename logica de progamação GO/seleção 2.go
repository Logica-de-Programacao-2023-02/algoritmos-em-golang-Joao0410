package main

import "fmt"

func main() {
	var num1, num2, num3 int

	// Solicitar os três números
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&num3)

	// Determinar o menor número
	menor := num1
	if num2 < menor {
		menor = num2
	}
	if num3 < menor {
		menor = num3
	}

	// Mostrar o menor número
	fmt.Printf("O menor número é: %d\n", menor)
}

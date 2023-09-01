package main

import "fmt"

func main() {
	var num1, num2 int

	// Solicitar os dois números
	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scan(&num2)

	// Verificar os números e realizar a operação apropriada
	if num1 >= 0 && num2 >= 0 {
		resultado := num1 * num2
		fmt.Printf("Ambos os números são positivos. Resultado da multiplicação: %d\n", resultado)
	} else {
		resultado := num1 + num2
		fmt.Printf("Pelo menos um dos números é negativo. Resultado da soma: %d\n", resultado)
	}
}

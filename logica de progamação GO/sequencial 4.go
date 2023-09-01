package main

import "fmt"

func main() {
	var n1, n2, n3, mediaPonderada float64

	// Pesos
	p1, p2, p3 := 2.0, 3.0, 5.0

	// Solicitando os números ao usuário
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&n1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&n2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&n3)

	// Calculando a média ponderada
	mediaPonderada = (n1*p1 + n2*p2 + n3*p3) / (p1 + p2 + p3)

	// Mostrando a média ponderada
	fmt.Printf("A média ponderada dos números com pesos 2, 3 e 5 é: %.2f\n", mediaPonderada)
}

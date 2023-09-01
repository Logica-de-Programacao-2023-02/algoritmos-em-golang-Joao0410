package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Println("Digite um número inteiro positivo:")
	fmt.Scan(&num)
	if num <= 0 {
		fmt.Println("Por favor, insira um número inteiro positivo válido.")
	}

	fmt.Printf("Divisores do número %d são:\n", num)
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}
}

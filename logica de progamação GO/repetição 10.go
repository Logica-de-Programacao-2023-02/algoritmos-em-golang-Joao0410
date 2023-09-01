package main

import (
	"fmt"
)

func main() {
	var num int
	var maior *int = nil

	fmt.Println("Digite números inteiros. Insira 0 para parar e mostrar o maior número:")

	for {
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		if maior == nil || num > *maior {
			maior = &num
		}
	}

	if maior == nil {
		fmt.Println("Nenhum número foi inserido!")
	} else {
		fmt.Printf("O maior número inserido é: %d\n", *maior)
	}
}

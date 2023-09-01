package main

import (
	"fmt"
)

func main() {
	var num, sum, count int

	fmt.Println("Digite números inteiros. Insira 0 para parar e calcular a média:")

	for {
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		sum += num
		count++
	}

	if count == 0 {
		fmt.Println("Nenhum número foi inserido!")
		return
	}

	media := float64(sum) / float64(count)
	fmt.Printf("A média aritmética dos números inseridos é: %f\n", media)
}

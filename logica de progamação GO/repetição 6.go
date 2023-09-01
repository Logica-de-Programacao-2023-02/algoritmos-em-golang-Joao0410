package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string

	fmt.Println("Digite um número para ver a sua tabuada:")
	fmt.Scan(&input)

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Por favor, insira um número válido.")
		os.Exit(1)
	}

	fmt.Printf("Tabuada de multiplicação do número %d:\n", num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}

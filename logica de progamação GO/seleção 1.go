package main

import "fmt"

func main() {
	var num1, num2 int

	//ler num
	fmt.Println("digite num1")
	fmt.Scan(&num1)

	fmt.Println("digite num2")
	fmt.Scan(&num2)

	//calcular
	if num1 > num2 {
		fmt.Print("O número 1 é maior de o Número 2")

	} else {

		fmt.Print("O número 1 é menor que o número 2 ")
	}
}

package main

import (
	"fmt"
)

func main() {
	var a, b, c float64

	// Solicitar os três números
	fmt.Print("Digite o primeiro número real: ")
	fmt.Scan(&a)

	fmt.Print("Digite o segundo número real: ")
	fmt.Scan(&b)

	fmt.Print("Digite o terceiro número real: ")
	fmt.Scan(&c)

	// Mostrar os números em ordem crescente
	if a < b && a < c {
		fmt.Print(a, ", ")
		if b < c {
			fmt.Print(b, ", ", c)
		} else {
			fmt.Print(c, ", ", b)
		}
	} else if b < a && b < c {
		fmt.Print(b, ", ")
		if a < c {
			fmt.Print(a, ", ", c)
		} else {
			fmt.Print(c, ", ", a)
		}
	} else {
		fmt.Print(c, ", ")
		if a < b {
			fmt.Print(a, ", ", b)
		} else {
			fmt.Print(b, ", ", a)
		}
	}
	fmt.Println()
}

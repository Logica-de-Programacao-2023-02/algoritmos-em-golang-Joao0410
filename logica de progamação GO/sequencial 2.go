package main

import "fmt"

func main() {
	var num int

	//ler num
	fmt.Println("digite num")
	fmt.Scan(&num)

	// Calculando e mostrando o dobro, triplo e quadruplo
	fmt.Println("Dobro:", num*2)
	fmt.Println("Triplo:", num*3)
	fmt.Println("Quadruplo:", num*4)
}

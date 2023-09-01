package main

import "fmt"

func main() {
	var idadeAnos float64

	// Solicitar a idade em anos
	fmt.Print("Digite sua idade em anos: ")
	fmt.Scan(&idadeAnos)

	// Calcular a idade em dias (considerando uma média de 365.25 dias por ano devido aos anos bissextos)
	idadeDias := idadeAnos * 365.25

	fmt.Printf("Sua idade aproximada em dias é: %.0f dias\n", idadeDias)

}

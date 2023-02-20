package main

import "fmt"

func main() {

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menotr ou igual a 15")
	}

	//variável "outroNumero" criada na própria condição
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else if numero < 10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número é menor do que zero")
	}
}

package main

import "fmt"

func main() {

	//FUNÇÃO ANÔNIMA NÃO TEM NOME

	// func(texto string) {
	// 	fmt.Println(texto)
	// }("Passando parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro")

	fmt.Println(retorno)
}

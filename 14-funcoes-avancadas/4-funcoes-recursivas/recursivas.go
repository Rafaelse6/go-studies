package main

import "fmt"

//FUNÇÕES RECURSIVAS SÃO FUNÇÕES QUE CHAMAM ELAS MESMAS, DIFICIL DE VER EM APLICAÇÕES REAIS

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	posicao := uint(12)

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

	fmt.Println(fibonacci(posicao))
}

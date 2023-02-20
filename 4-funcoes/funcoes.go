package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicso(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função f")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicso(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
	/*
		resultadoSoma, _ := calculosMatematicso(10, 15)
		fmt.Println(resultadoSoma)
		O "_" SERVE PARA AINDA FAZER O QUE A FUNÇÃO DETERMINA, MAS NÃO MOSTRA O RESULTADO
	*/
}

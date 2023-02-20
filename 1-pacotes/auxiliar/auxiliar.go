package auxiliar

import "fmt"

//Funcao com letra minuscula só vai ser visível dentro do pacote em que está(private) e se for com letra maiúscula fica como public(pode ser exportada por outros pacotes)
func Escrever() {
	fmt.Println("Escrevendo do arquivo auxiliar")
	escrever2()
}

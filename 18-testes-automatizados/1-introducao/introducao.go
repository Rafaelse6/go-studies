package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Praça das Rosas")
	fmt.Println(tipoEndereco)
}

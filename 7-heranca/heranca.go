package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

//PARA "HERANÇA" é só não passar o tipo, como abaixo
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Print(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome)
}

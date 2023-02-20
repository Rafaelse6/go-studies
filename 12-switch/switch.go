package main

import "fmt"

func diaDaSemana(numero int) string {
	var diaDaSemana string
	switch numero {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough
	case 2:
		diaDaSemana = "Segunda"
	case 3:
		diaDaSemana = "Terça"
	case 4:
		diaDaSemana = "Quarta"
	case 5:
		diaDaSemana = "Quinta"
	case 6:
		diaDaSemana = "Sexta"
	case 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	default:
		return "IDK"
	}
}

func main() {
	dia := diaDaSemana(10)
	fmt.Println(dia)

	fmt.Println("-------")
	dia2 := diaDaSemana(1)
	fmt.Println(dia2)
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -10000000
	fmt.Println(numero)

	//uint = só pode ser número 0 ou números positivos
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias(apelidos de tipos(rune, byte, etc))
	//INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//NÚMEROS REAIS

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	//FIM DOS NÚMEROS REAIS

	//STRING
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)
	//NÃO EXISTE CHAR EM GO

	//FIM STRINGS

	var texto int16
	fmt.Println(texto)

	//BOOLEAN

	var booleano1 bool = true
	fmt.Println(booleano1)

	//FIM BOOLEAN

	//ERROR
	var error error = errors.New("Erro interno")
	fmt.Println(error)
}

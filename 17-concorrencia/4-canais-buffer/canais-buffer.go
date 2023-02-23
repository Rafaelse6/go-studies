package main

import "fmt"

func main() {

	//Canal com buffer só bloqueia quando atingir a capacidade máxima

	canal := make(chan string, 2)
	canal <- "Olá mundo!"
	canal <- "Programanado em Go!"
	//canal <- "Programanado em Go!" = deadlock

	mensagem := <-canal
	mensagem2 := <-canal
	//mensagem3 := <-canal = deadlock

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}

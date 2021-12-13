package main

import (
	"fmt"
)

func main() {

	// var aprovados map[int]string
	// mapas devem ser inicializados

	aprovados := make(map[int]string) // chave do map do tipo inteira e o valor do map tipo string

	aprovados[12345678978] = "Maria" //Exemplo: o CPF é minha chave inteira e o nome da pessoa é minha string
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d\n", nome, cpf)

	}
}

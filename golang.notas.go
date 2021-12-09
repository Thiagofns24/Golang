package main

import (
	"fmt"
)

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"

	}

	return "Reprovado" //Float64 é uma versão float que armazena valores decimais usando um total de 64 bits de dados
}

func main() {

	fmt.Println(obterResultado(2.2))

}

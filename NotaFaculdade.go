package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough //Programa entende que tanto a nota 10, quanto a nota 9 vão geral o conceito A
	case 9:
		return "A"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D" //fallthrough
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida" //Se a palavra-chave fallthrough estiver presente no bloco de caso,
		//ele transferirá o controle para o próximo caso,
		//mesmo que o caso atual possa ter correspondido.

	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(-2.1))

}

// No exemplo [anterior](range), vimos como `for` e
// `range` fornecem uma iteração sobre as estruturas
// básicas de dados. Também podemos usar essa sintaxe
// para iterar sobre os valores recebidos de um canal.

package main

import "fmt"

func main() {

	// Vamos iteraragir com mais de dois valores na
	// fila do canal.
	queue := make(chan string, 2)
	queue <- "um"
	queue <- "dois"
	close(queue)

	// Esse `range` iterage sobre cada elemento à medida
	// que é recebido da fila `queue`. Como fechamos o
	// canal acima, a iteração termina depois de receber
	// os dois elementos.
	for elem := range queue {
		fmt.Println(elem)
	}
}

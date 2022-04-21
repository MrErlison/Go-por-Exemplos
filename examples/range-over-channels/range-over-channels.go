// No exemplo [anterior](range), vimos como `for` e 
// `range` fornecem uma iteração sobre as estruturas 
// básicas de dados. Também podemos usar essa sintaxe 
// para iterar sobre os valores recebidos de um canal.
package main

import "fmt"

func main() {

	// We'll iterate over 2 values in the `queue` channel.
	// Vamos iterar mais de 2 valores na file do canal.
	queue := make(chan string, 2)
	queue <- "um"
	queue <- "dois"
	close(queue)

	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.
	// Esse intervalo itera sobre cada elemento à medida que é recebido da fila. Como fechamos o canal acima, a iteração termina depois de receber os 2 elementos.
	for elem := range queue {
		fmt.Println(elem)
	}
}

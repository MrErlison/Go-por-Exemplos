// O Go tem suporte para múltiplos valores de retorno. 
// Esse recurso é usado com frequência no Go idiomático,
// por exemplo, para retornar valores de resultado e 
// erro de uma função.

package main

import "fmt"

// O `(int, int)` nessa função mostra que a função 
// retorna dois valores do tipo `int`.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Aqui usamos os dois valores de retorno diferentes
	// da chamada com atribuição múltipla
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Se você quiser apenas um elemento dos valores 
	// retornados, use o identificador em branco `_`
	_, c := vals()
	fmt.Println(c)
}

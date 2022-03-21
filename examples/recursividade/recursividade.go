// Go suporta
// <a href="https://pt.wikipedia.org/wiki/Recursividade_(ciência_da_computação)"><em>funções recursivas</em></a>.
// Aqui temos um exemplo clássico.

package main

import "fmt"

// Essa função `fact` chama a sim mesma até chegar ao
// caso base do `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// _Closures_ também podem ser recursivos, mas isso requer que
	// o _closure_ seja declarado explicitamente com um tipo `var`
	// antes de ser definido.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Como a `fib` foi declarada anteriormente na 
		// função principal, `main`, o Go sabe qual 
		// função chamar com a `fib` aqui.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

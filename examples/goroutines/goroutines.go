// Uma _goroutine_ é uma _thread_ leve de execução.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suponha que tenhamos uma chamada de função `f(s)`.
	// Veja como chamaríamos isso da maneira usual, 
	// executando-o de forma síncrona.
	f("direct")

	// Para invocar essa função em uma _goroutine_,
	// use `go f(s)`. Esta nova _goroutine_ será executada
	// simultaneamente com a chamada.
	go f("goroutine")

	// Você também pode iniciar uma rotina para uma 
	// chamada de função anônima.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Nossas duas chamadas de função estão sendo executadas de
	// forma assíncrona em _goroutines_ separadas agora. Espere
	// que eles terminem (para uma abordagem mais robusta, use 
	// um [WaitGroup](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("done")
}
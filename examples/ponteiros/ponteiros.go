// O Go suporta <em><a href="https://pt.wikipedia.org/wiki/Ponteiro_(programação)">ponteiros</a></em>, 
// permitindo que você passe referências a valores e
// registros dentro do seu programa

package main

import "fmt"

// Mostraremos como os ponteiros funcionam em contraste
// aos valores com duas funções: `zeroval` e `zeroptr`.
// A `zeroval` tem um parâmetro `int`, então os 
// argumentos serão passados para ele por valor. A
// `zeroval` receberá uma cópia do `ival` distinta da 
// função de chamada.
func zeroval(ival int) {
	ival = 0
}

// A `zeroptr`, por outro lado, tem um parâmetro `*int`, 
// o que significa que é preciso um ponteiro `int`. O 
// código `*iptr` no corpo da função _desreferencia_ 
// o ponteiro de seu endereço de memória para o valor
// atual nesse endereço. Atribuir um valor a um ponteiro
// desreferenciado altera o valor no endereço 
// referenciado.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// A sintaxe `&i` fornece o endereço de memória de `i`, 
	// ou seja, um ponteiro para `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Os ponteiros também podem ser impressos.
	fmt.Println("pointer:", &i)
}

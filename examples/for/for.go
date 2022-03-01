// `for` é o único construtor de laços da linguagem Go.
// Aqui estão alguns tipos básicos de laço `for`.

package main

import "fmt"

func main() {

	//  O tipo mais comum, com uma única condição.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Um clássico laço `for` inicio/condição/final
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` sem uma condição ficará em repetição até
	// o uso do `break` ou `return` para finalizar.
	for {
		fmt.Println("loop")
		break
	}

	// Você pode usar `continue` para ir direto para a
	// próxima interação do laço.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

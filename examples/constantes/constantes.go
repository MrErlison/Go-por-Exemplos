// Go suporta _constantes_ de caracteres, string, boolean,
// e valores numéricos

package main

import (
	"fmt"
	"math"
)

// `const` declara um valor constante.
const s string = "constante"

func main() {
	fmt.Println(s)

	// Uma instrução `const` pode aparecer em
	// qualquer lugar, tal como a instrução `var`
	const n = 500000000

	// Expressões constantes realizam aritimética
	// com precisão arbitrária.
	const d = 3e20 / n
	fmt.Println(d)

	// Uma constante numérica não tem nenhum tipo
	// até a atribuição de um, tal como uma conversão
	// explícita.
	fmt.Println(int64(d))

	// Um número pode receber um tipo usando-o
	// em um contexto que requer um, como uma
	// atribuição de variável ou chamada de função.
	// Por exemplo, aqui o `math.Sin` espera um
	// `float64`.
	fmt.Println(math.Sin(n))
}

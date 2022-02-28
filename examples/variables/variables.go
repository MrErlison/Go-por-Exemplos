// Em Go, _variáveis_ são explicitamente declaradas e
// usadas pelo compilador. Por exemplo, verificar
// type-correctness de chamada de funções.

package main

import "fmt"

func main() {

	// `var` declara uma ou mais variáveis.
	var a = "initial"
	fmt.Println(a)

	// Você pode declarar múltiplas variáveis de uma vez só.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go inferirá o tipo de variável inicializada.
	var d = true
	fmt.Println(d)

	// As variáveis declaradas sem uma inicialização
	// correspondente terão _valor zero_. Por exemplo,
	// o valor zero para um `int` é `0`.

	var e int
	fmt.Println(e)

	// A sintaxe `:=` é uma abreviação para declarar
	// e inicializar uma variável. Por exemplo,
	// `var f string = "apple"`.
	f := "apple"
	fmt.Println(f)
}

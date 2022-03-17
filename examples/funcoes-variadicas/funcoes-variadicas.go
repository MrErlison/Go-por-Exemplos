// [_Funções Variádicas_](http://en.wikipedia.org/wiki/Variadic_function)
// podem ser chamadas com qualquer número de argumentos
// à direita. Por exemplo, `fmt.Println` é uma função
// variádica comum.

package main

import "fmt"

// Aqui está uma função que terá um número arbitrário de
// inteiros como argumentos.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// As funções variádicas podem ser chamadas da 
	// maneira usual com argumentos individuais.
	sum(1, 2)
	sum(1, 2, 3)

	// Se você já tem vários argumentos em uma _slice_,
	// aplique-os a uma função variádica desse modo
	// _func(slice...)_.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
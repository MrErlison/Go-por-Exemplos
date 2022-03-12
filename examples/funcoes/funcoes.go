// As funções são centrais no Go. Aprenderemos sobre as
// funções com alguns exemplos diferentes.

package main

import "fmt"

// Aqui está uma função que pega dois `int` e retorna
// sua soma como um inteiro.
func plus(a int, b int) int {

	// Go requer retornos explícitos, ou seja, não 
	// retornará automaticamente o valor da última
	// expressão.
	return a + b
}

// Quando você tem vários parâmetros consecutivos do
// mesmo tipo, você pode omitir o nome do tipo para
// os parâmetros com tipo semelhante até o parâmetro
// final que declara o tipo.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Chame uma função exatamente como esperado, com 
	// nome(argumentos).
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

// Em Go, um _array_ é uma sequência numerada de 
// elementos de um comprimento específico.

package main

import "fmt"

func main() {

	// Aqui criamos um _array_ `a` que conterá 
	// exatamente 5 elementos do tipo `int`. 
	// O tipo de elementos e o comprimento fazem parte 
	// do tipo do _array_. Por padrão, um _array_ tem
	// valor zero, o que para `int` significa `0`s.
	var a [5]int
	fmt.Println("emp:", a)

	// Nós podemos definir um valor em um índice usando
	// a sintaxe `array[índice] = valor` e obter um valor
	// com `array[índice]`.
	a[4] = 100
	fmt.Println("a:", a)
	fmt.Println("a[4]:", a[4])

	// A função `len` retorna o comprimento de um _array_.
	fmt.Println("comprimento:", len(a))

	// Use essa sintaxe para declarar e inicializar 
	// um _array_ em uma linha.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// Os tipos de _array_ são unidimensionais, mas 
	// você pode compor os tipos para construir 
	// estruturas de dados multidimensionais.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

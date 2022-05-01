// A partir da versão 1.18, a linguagem Go adicionou 
// suporte para _genéricos_, também conhecidos como 
// _parâmetros de tipo_.

package main

import "fmt"

// Como exemplo de uma função genérica, o `MapKeys` pega
// um `map` de qualquer tipo e retorna um `slice` de 
// chaves. Esta função tem dois parâmetros de tipo - `K`
// e `V`; `K` é comparável a uma constante, o que 
// significa que podemos comparar valores desse tipo com
// os operadores `==` e `!=`. Isso é necessário para o 
// `map` de chaves em Go. `V` tem a constante `any`, o 
// que significa que não é restrito de forma alguma 
// (`any` é um alias para `interface{}`).
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Como exemplo de um tipo _genérico_, `List` é uma lista
// vinculada individualmente com valores de qualquer tipo.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Podemos definir métodos em tipos genéricos, assim como
// fazemos em tipos regulares, mas temos que manter os
// parâmetros de tipo no lugar. O tipo é `List[T]`, 
// não `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// Ao invocar as funções genéricas, muitas vezes 
	// podemos confiar na inferência de tipos. Observe
	// que não precisamos especificar os tipos para `K`
	// e `V` ao chamar `MapKeys` - o compilador os
	// infere automaticamente.
	fmt.Println("chaves m:", MapKeys(m))

	// ... embora também possamos especificá-los
	// explicitamente.
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("lista:", lst.GetAll())
}


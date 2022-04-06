// No Go, é idiomático comunicar erros por meio de um
// valor de retorno explícito e separado. Isso contrasta
// com as exceções usadas em linguagens como Java e Ruby
// em que o valor sobrecarregado de resultado é o único
// erro às vezes usado em C. A abordagem do Go facilita
// a visualização de quais funções retornam erros e o 
// manuseio deles usando as mesmas construções de 
// linguagem empregadas para quaisquer outras tarefas sem
// erros.

package main

import (
	"errors"
	"fmt"
)

// Por convenção, os erros são os últimos valores 
// retornados e têm um tipo `error`, que é uma interface
// interna.
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` constrói um valor básico de erro
		// com a mensagem de erro fornecida.
		return -1, errors.New("Não consigo trabalhar com 42")

	}

	// Um valor nulo na posição de erro indica que não
	// houve erro.
	return arg + 3, nil
}

// É possível usar tipos personalizados de erros 
// implementando o método `Error()`. Aqui está uma 
// variante no exemplo acima que usa um tipo 
// personalizado para representar explicitamente um
// erro de argumento.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// Nesse caso, usamos a sintaxe `&argError` para 
		// construir uma nova estrutura, fornecendo 
		// valores para os dois campos `arg` e `prob`.
		return -1, &argError{arg, "Não consigo trabalhar com isso"}
	}
	return arg + 3, nil
}

func main() {

	// Os dois loops abaixo testam cada uma das nossas
	// funções de retorno de erros. Observe que o uso de
	// uma verificação de erro no `if` é uma expressão 
	// idiomática comum no código Go.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 falhou:", e)
		} else {
			fmt.Println("f1 funcionou:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 falhou:", e)
		} else {
			fmt.Println("f2 funcionou:", r)
		}
	}

	// Se você quiser usar programaticamente os dados em
	// um erro personalizado, precisará obter o erro como
	// uma instância do tipo de erro personalizado por meio
	// de asserção de tipo.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
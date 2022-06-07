// O Go possibilita a recuperação de um _panic_, 
// usando a função `recover` integrada. Uma função
// `recover` pode impedir que um `panic` aborte o programa
// para continuar com a execução.

// Um exemplo de onde isso pode ser útil: um servidor 
// não gostaria de falhar se uma das conexões do cliente
// apresentar um erro crítico. Em vez disso, o servidor
// gostaria de fechar essa conexão e continuar atendendo
// a outros clientes. Na verdade, é isso que a rede/http
// do Go faz por padrão para servidores HTTP.

package main

import "fmt"

func mayPanic() {
	panic("um problema")
}

func main() {
	// A função `recover` deve ser chamada dentro de uma
	// função `defer`. Quando a função de fechamento 
	// entrar em `panic`, o adiamento será ativado e uma 
	// chamada de `recover` dentro dele pegará o `panic`.
	defer func() {
		if r := recover(); r != nil {
			// O valor de retorno do `recover` é o erro gerado na
			// chamada `panic`.
			fmt.Println("Recuperado. Erro:\n", r)
		}
	}()

	mayPanic()

	// Esse código não será executado, porque `mayPanic`
	// entra em `panic`. A execução das principais 
	// paradas no ponto do `panic` é retomado na 
	// função `defer`.
	fmt.Println("Depois mayPanic()")
}

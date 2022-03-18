// O Go suporta [funções anônimas](http://en.wikipedia.org/wiki/Anonymous_function),
// que podem formar <a href="http://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Funções anônimas são úteis quando você deseja definir 
// uma função em linha sem ter que nomeá-la.

package main

import "fmt"

// Esta função `intSeq` retorna outra função, que 
// definimos anonimamente no corpo da função `intSeq`.
// A função retornada se fecha sobre a variável `i` para 
// formar um _closure_.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Chamamos `intSeq`, atribuindo o resultado 
	// (uma função) ao `nextInt`. Este valor de função 
	// captura seu próprio valor `i`, que será atualizado
	// toda vez que chamarmos `nextInt`.
	nextInt := intSeq()

	// Veja o efeito do _closure_ ligando para `nextInt`
	// algumas vezes.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Para confirmar que o estado é exclusivo dessa
	// função específica, crie e teste uma nova.
	newInts := intSeq()
	fmt.Println(newInts())
}
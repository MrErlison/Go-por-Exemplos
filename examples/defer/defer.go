// O _defer_ é usado para garantir que uma chamada de 
// função seja executada mais tarde na execução de um
// programa, geralmente para fins de limpeza. O `defer`
// é frequentemente usado onde, por exemplo, `ensure` e
// `finally`, seria usado em outras linguagens.

package main

import (
	"fmt"
	"os"
)

// Suponha que criássemos um arquivo, para escrever
// nele e para depois fecharmos. Veja como poderíamos
// fazer isso com o `defer`.
func main() {

	// Imediatamente após obter um objeto de arquivo com
	// `createFile`, adiamos o fechamento desse arquivo
	// com `closeFile`. Isso será executado no final da
	// função de principal, após o término 
	// do `writeFile`.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("criando")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("escrevendo")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("fechando")
	err := f.Close()
	// É importante verificar se há erros ao fechar um
	// arquivo, mesmo em uma função adiada.
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

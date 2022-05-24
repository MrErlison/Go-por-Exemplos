// Escrever arquivos no Go segue padrões semelhantes aos
// que vimos anteriormente para leitura.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Para começar, veja como despejar uma string (ou
	// apenas bytes) em um arquivo.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// Para gravações mais granulares, abra um arquivo 
	// para escrever.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// É comum adiar o `Close` imediatamente após
	// a abertura de um arquivo.
	defer f.Close()

	// Você pode escrever `slices` de bytes.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// Uma `WriteString` também está disponível
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Chame o `Sync` para um armazenamento estável.
	f.Sync()

	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Use `Flush` para garantir que todas as operações
	// armazenadas em buffer tenham sido aplicadas.
	w.Flush()

}

// Uma string Go é um _slices_ de bytes somente de 
// leitura. A linguagem e a biblioteca padrão tratam
// strings especialmente - como contêineres de texto
// codificados em [UTF-8](https://pt.wikipedia.org/wiki/UTF-8). 
// Em outros idiomas, as strings são feitas de 
// "caracteres". Em Go, o conceito de um caractere é
// chamado de `rune` - é um inteiro que representa um 
// ponto de código Unicode. Esta postagem do [blog Go](https://go.dev/blog/strings)
// é uma boa introdução ao tópico.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` é uma `string` que atribui um valor literal 
	// que representa a palavra "olá" no idioma 
	// tailandês. As strings literais de Go são textos
	// codificados em UTF-8.
	const s = "สวัสดี"

	// Como as strings são equivalentes a `[]byte`, isso
	// produzirá o comprimento dos bytes brutos 
	// armazenados.
	fmt.Println("Len:", len(s))

	// A indexação em uma string produz os valores brutos
	// de bytes em cada índice. Este loop gera os valores
	// hexadecimais de todos os bytes que constituem os
	// pontos de código em `s`.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Para contar quantas _runes_ existentes em uma 
	// string, podemos usar o pacote `utf8`. Observe que
	// o tempo de execução do `RuneCountInString` depende
	// do tamanho da string, porque ele tem que 
	// decodificar cada rune UTF-8 sequencialmente. 
	// Alguns caracteres tailandeses são representados 
	// por vários pontos de código UTF-8, então o 
	// resultado dessa contagem pode ser surpreendente.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// Um loop `range` lida especialmente com strings
	// e decodifica cada `rune` junto com seu 
	// deslocamento na string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Podemos obter a mesma iteração usando a função 
	// `utf8.DecodeRuneInString` explicitamente.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// Isso demonstra a passagem do `rune` para uma função.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// Os valores incluídos entre aspas simples são 
	// _rune_ literais. Podemos comparar um valor de 
	// _rune_ diretamente com um _rune_ literal.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

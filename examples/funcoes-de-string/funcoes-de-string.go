// A biblioteca padrão de strings  fornece muitas funções
// úteis relacionadas a strings. Aqui estão alguns 
// exemplos para lhe dar uma noção do pacote.

package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	// Aqui está um exemplo das funções disponíveis em `strings`.
	// Como essas são funções do pacote, nós precisamos passar a 
	// string em questão como o primeiro argumento para a função.
	// Você pode encontrar mais funções nos documentos do pacote
	// strings.
	p("Contem:    ", s.Contains("test", "es"))
	p("Conta:     ", s.Count("test", "t"))
	p("É Prefixo: ", s.HasPrefix("test", "te"))
	p("É Sufixo:  ", s.HasSuffix("test", "st"))
	p("Índice:    ", s.Index("test", "e"))
	p("Concatena: ", s.Join([]string{"a", "b"}, "-"))
	p("Repete:    ", s.Repeat("a", 5))
	p("Substitui: ", s.Replace("foo", "o", "0", -1))
	p("Substitui: ", s.Replace("foo", "o", "0", 1))
	p("Separa:    ", s.Split("a-b-c-d-e", "-"))
	p("Minúscula: ", s.ToLower("TEST"))
	p("Maiúscula: ", s.ToUpper("test"))
}

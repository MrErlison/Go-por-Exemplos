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
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}

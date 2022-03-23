// As estruturas do Go são tipos de coleções de campos.
// Elas são úteis para agrupar dados para formar registros.

package main

import "fmt"

// Este tipo de estrutura de `pessoa` tem os campos 
// `nome` e `idade`.
type pessoa struct {
	nome string
	idade  int
}

// `novaPessoa` constrói uma nova estrutura de pessoa com o nome próprio.
func novaPessoa(nome string) *pessoa {
	// Você pode retornar com segurança um ponteiro para
	// a variável local, pois uma variável local 
	// sobreviverá ao escopo da função.
	p := pessoa{nome: nome}
	p.idade = 42
	return &p
}

func main() {

	// Esta sintaxe cria uma nova estrutura.
	fmt.Println(pessoa{"Pedro", 20})

	// Você pode nomear os campos ao inicializar uma estrutura.
	fmt.Println(pessoa{nome: "Alice", idade: 30})

	// Os campos omitidos terão valor zero.
	fmt.Println(pessoa{nome: "Frederico"})

	// Um prefixo `&` produz um ponteiro para a estrutura.
	fmt.Println(&pessoa{nome: "Ana", idade: 40})

	// É idiomático encapsular a criação de novas estruturas em funções
	fmt.Println(novaPessoa("Tiago"))

	// Acesse campos de estrutura com um ponto.
	s := pessoa{nome: "Antonio", idade: 50}
	fmt.Println(s.nome)

	// Você também pode usar pontos com ponteiros de
	// estrutura - os ponteiros são automaticamente 
	// desreferenciados.
	sp := &s
	fmt.Println(sp.idade)

	// As estruturas são mutáveis.
	sp.idade = 51
	fmt.Println(sp.idade)
}

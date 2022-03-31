// Go suporta _métodos_ definidos em tipos de estrutura..

package main

import "fmt"

type retangulo struct {
	largura, altura int
}

// This `area` method has a _receiver type_ of 
// `*retangulo`. Este método de `area` tem um _tipo de 
// receptor_ de `*retangulo`.
func (r *retangulo) area() int {
	return r.largura * r.altura
}

// Os métodos podem ser definidos para receptores do
// tipo de ponteiro ou de valor. Aqui está um exemplo
// de um receptor de valor.
func (r retangulo) perimetro() int {
	return 2*r.largura + 2*r.altura
}

func main() {
	r := retangulo{largura: 10, altura: 5}

	// Aqui chamamos os dois métodos definidos para a
	// nossa estrutura.
	fmt.Println("area: ", r.area())
	fmt.Println("perimetro:", r.perimetro())

	// Go lida automaticamente com a conversão entre
	// valores e ponteiros para chamadas de método. Você
	// pode querer usar um tipo de receptor de ponteiro
	// para evitar copiar chamadas de método ou para
	// permitir que o método mude a estrutura receptora.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimetro:", rp.perimetro())
}

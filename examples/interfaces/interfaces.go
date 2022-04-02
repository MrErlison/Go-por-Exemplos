// _Interfaces_ são chamadas de coleções de assinaturas
// de métodos

package main

import (
	"fmt"
	"math"
)

// Aqui está uma interface básica para formas geométricas
type geometria interface {
	area() float64
	perimetro() float64
}

// Para o nosso exemplo, implementaremos essa interface
// nos tipos de `retangulo` e `circulo`.
type retangulo struct {
	largura, altura float64
}
type circulo struct {
	raio float64
}

// Para implementar uma interface no Go, só precisamos
// implementar todos os métodos na interface. Aqui
// implementamos `geometria` em `retangulos`.
func (r retangulo) area() float64 {
	return r.largura * r.altura
}
func (r retangulo) perimetro() float64 {
	return 2*r.largura + 2*r.altura
}

// A implementação para `circulos`.
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// Se uma variável tiver um tipo de interface, podemos 
// chamar os métodos que estão na interface nomeada. Aqui
// está uma função genérica de `medida` que aproveita 
// isso para trabalhar em qualquer `geometria`.
func medida(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	r := retangulo{largura: 3, altura: 4}
	c := circulo{raio: 5}

	// Os tipos de estrutura `circulo` e `retangulos` implementam
	// a interface de `geometria` para que possamos usar instâncias
	// dessas estruturas como argumentos para `medida`.
	medida(r)
	medida(c)
}

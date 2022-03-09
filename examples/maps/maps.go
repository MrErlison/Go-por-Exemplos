// _Map_ é o tipo de [dados associativos internos do Go](http://en.wikipedia.org/wiki/Associative_array)
// (às vezes chamados de _hashes_ ou _dicts_ em outras linguagens).

package main

import "fmt"

func main() {

	// Para criar um _map_ vazio, use o _make_:
	// `make(map[key-type]val-type)`.
	m := make(map[string]int)

	// Defina chave/valor usando a sintaxe 
	// típica `name[key] = val`
	m["k1"] = 7
	m["k2"] = 13

	// Imprimir um _map_ com _fmt.Println_ mostrará 
	// todos os seus pares chave/valor.
	fmt.Println("map:", m)

	// Obtenha um valor para uma chave com
	// `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// O _len_ retorna o número de pares chave/valor
	// quando chamado em um _map_.
	fmt.Println("len:", len(m))

	// O _delete_ remove os pares de chave/valor
	// de um _map_.
	delete(m, "k2")
	fmt.Println("map:", m)

	// O segundo valor é retornado ao obter um
	// valor de um _map_ indicando se a chave esta 
	// presente no _map_. Isso pode ser usado para 
	// distinguir entre chaves ausentes e chaves com 
	// valores zero como 0 ou "". Aqui não precisávamos
	// do valor em si, então o ignoramos com o
	// identificador em branco _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Você também pode declarar e inicializar um novo
	// mapa na mesma linha com essa sintaxe.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

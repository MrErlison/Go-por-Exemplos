// _range_ iterage sobre elementos em uma variedade de 
// estruturas de dados. Vamos ver como usar o range com
// algumas das estruturas de dados que já aprendemos.

package main

import "fmt"

func main() {

	// Aqui usamos _range_ para somar os números em um _slice_. 
	// _Arrays_ funcionam da mesma forma.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// _range_ nos _arrays_ e nos _slices_ fornece o
	// índice e o valor para cada entrada. Acima, não
	// precisávamos do índice, então o ignoramos com o
	// identificador em branco _. Às vezes, na verdade, 
	// queremos os índices.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// _range_ nos _map_ itera sobre o par chaves/valores.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// _range_ também pode iterar apenas sobre as chaves 
	// de um _map_.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// _range_ nas _strings_ itera sobre os pontos de 
	// código Unicode. O primeiro valor é o índice do 
	// bytes inicial da _rune_ e o segundo a própria 
	// _run_. Consulte [Strings e Runes](strings-and-runes)
	// para mais detalhes.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

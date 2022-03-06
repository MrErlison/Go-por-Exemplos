// _Slices_ são um tipo de dados no Go, dando uma
// `interface` mais poderosa para sequências do 
// que `array`.

package main

import "fmt"

func main() {

	// Ao contrário dos `arrays`, os `slices` tem seus
	// tipos apenas pelos elementos que contêm (não pelo
	// número de elementos). Para criar um `slice` vazio
	// com comprimento diferente de zero, use construtor
	// `make`. Aqui fazemos os `slices` de `string` de 
	// comprimento 3 (inicialmente com valor zero).

	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Podemos definir e obter exatamente como o `array`.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("seet", s)
	fmt.Println("get:", s[2])

	// `len` retorna o tamanho do `slice`.
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.
	
	// Além dessas operações básicas, os `slices` 
	// suportam muito mais, o que os tornam mais
	// rico do que o `array`. 
	// Um deles é o `append`, que retorna um `slice` 
	// contendo um ou mais novos valores. Observe que
	// precisamos aceitar um valor de retorno do 
	// `append`, bem como podemos obter um novo valor
	// do `slice`.

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// `Sllices` também podem ser copiadas. Aqui criamos
	// um `slice` vazio `c` do mesmo comprimento que `s` 
	// e copiamos para `c` o `s`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// `Slices` suportam um operador “slice” com a 
	// sintaxe `slice[low:high]`. Por exemplo, isso obtém
	// um `slice` dos elementos s[2], s[3] e s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Esse `slice` corta até (excluindo) `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l)

	// E esse `slice` parte de (incluindo) `s[2]`.
	l = s[2:]
	fmt.Println("sl3:", l)

	// Também podemos declarar e inicializar uma variável
	// `slice` em uma única linha.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// `Slices` podem ser compostos em estruturas de 
	// dados multidimensionais. O comprimento dos 
	// `slices` internos pode variar, ao contrário dos
	// `array` multidimensionais.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

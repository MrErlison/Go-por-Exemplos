// _Slices_ são um tipo de dados no Go, dando uma
// _interface_ mais poderosa para sequências do 
// que _array_.

package main

import "fmt"

func main() {

	// Ao contrário dos _arrays_, os _slices_ tem seus
	// tipos apenas pelos elementos que contêm (não pelo
	// número de elementos). Para criar um _slice_ vazio
	// com comprimento diferente de zero, use construtor
	// _make_. Aqui fazemos os _slices_ de _string_ de 
	// comprimento 3 (inicialmente com valor zero).
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Podemos definir e obter exatamente como o _array_.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("seet", s)
	fmt.Println("get:", s[2])

	// _len_ retorna o tamanho do _slice_.
	fmt.Println("len:", len(s))

	// Além dessas operações básicas, os _slices_
	// suportam muito mais, o que os tornam mais
	// rico do que o _array_. 

	// Um deles é o _append_, que retorna um _slice_
	// contendo um ou mais novos valores. Observe que
	// precisamos aceitar um valor de retorno do 
	// _append_, bem como podemos obter um novo valor
	// do _slice_.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// _Slices_ também podem ser copiadas. Aqui criamos
	// um _slice_ vazio _c_ do mesmo comprimento que _s_
	// e copiamos para _c_ o _s_.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// _Slices_ suportam um operador _slice_ com a 
	// sintaxe _slice[low:high]_. Por exemplo, isso obtém
	// um _slice_ dos elementos s[2], s[3] e s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Esse _slice_ corta até (excluindo) _s[5]_.
	l = s[:5]
	fmt.Println("sl2:", l)

	// E esse _slice_ parte de (incluindo) _s[2]_.
	l = s[2:]
	fmt.Println("sl3:", l)

	// Também podemos declarar e inicializar uma variável
	// _slice_ em uma única linha.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// _Slices_ podem ser compostos em estruturas de 
	// dados multidimensionais. O comprimento dos 
	// _slices_ internos pode variar, ao contrário dos
	// _array_ multidimensionais.
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

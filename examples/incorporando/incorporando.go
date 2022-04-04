// Go apoia a incorporação de estruturas e interfaces 
// para expressar uma composição de tipos.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// O tipo `contêiner` incorpora uma estrutura `base`. A
// incorporação parece um campo sem nome.
type container struct {
	base
	str string
}

func main() {

	// Ao criar estruturas com literais, temos que 
	// inicializar a incorporação explicitamente; aqui
	// o tipo incorporado serve como o nome do campo.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Podemos acessar os campos da base diretamente 
	// em `co`, por exemplo, `co.num`
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Como alternativa, podemos adicionar o caminho 
	// completo usando o nome do tipo incorporado.
	fmt.Println("also num:", co.base.num)

	// Como o `contêiner` incorpora a `base`, os métodos
	// da `base` também se tornam métodos do `contêiner`.
	// Aqui invocamos um método que foi incorporado a 
	// partir da `base` diretamente no `co`.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// A incorporação de estruturas com métodos pode ser
	// usada para conceder implementações de interface a
	// outras estruturas. Aqui vemos que um `contêiner` 
	// agora implementa a interface do `descritor` porque 
	// incorpora a `base`.
	var d describer = co
	fmt.Println("describer:", d.describe())
}

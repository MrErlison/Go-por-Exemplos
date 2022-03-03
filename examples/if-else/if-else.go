// As ramificações com `if` e `else` em Go são diretas.

package main

import "fmt"

func main() {

	// Aqui está um exemplo básico.
	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é ímpar")
	}

	// Você pode ter uma instrução `if` sem o `else`
	if 8%4 == 0 {
		fmt.Println("8 é divisível por 4")
	}

	// Uma instrução pode preceder condicionais;
	// quaisquer variáveis declaradas nessa instrução
	// estão disponíveis em todas as ramificações.
	if num := 9; num < 0 {
		fmt.Println(num, "é negativo")
	} else if num < 10 {
		fmt.Println(num, "tem um dígito")
	} else {
		fmt.Println(num, "tem múltiplos dígitos")
	}
}

// Observe que você não precisa de parênteses para as
// condições em GO, mas que as chaves são necessárias.

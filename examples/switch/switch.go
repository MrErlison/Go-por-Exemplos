// As instruções _switch_ expressam condições através
// de diversas ramificações.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Aqui tem um `switch` básico.
	i := 2
	fmt.Print("Escreva ", i, " com ")
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	// Você pode usar vírgulas para separar múltiplas
	// expressões na mesma instrução `case`. Nós usamos
	// o `case` opcional `default` neste exemplo.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// `switch` sem uma expressão é um caminho 
	// alternativo para expressar a lógica `if/else`.
	// Aqui nós também mostramos como as expressões
	// `case` podem ser não constantes.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("É antes do meio dia")
	default:
		fmt.Println("É depois do meio dia")
	}

	// Um tipo `switch` compara tipos ao invés de valores
	// Você pode usar isso para descobrir o tipo de valor
	// de uma interface. No exemplo, a variável `t` terá
	// o tipo correspondente à sua cláusula.

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Eu sou um bool")
		case int:
			fmt.Println("Eu sou um int")
		default:
			fmt.Printf("Eu não sou do tipo %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Oi")
}

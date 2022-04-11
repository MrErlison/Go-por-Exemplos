// Por padrão, os canais não são armazenados em _buffer_,
// o que significa que eles só aceitarão envios 
// (`chan <-`) se houver um recebimento correspondente
// (`<- chan`) pronto para receber o valor enviado. 
// Os canais em buffer aceitam um número limitado de 
// valores sem um receptor correspondente para esses 
// valores.

package main

import "fmt"

func main() {

	// Aqui fazemos `make` um canal de `strings` em 
	// buffer de até 2 valores.
	messages := make(chan string, 2)

	// Como esse canal está armazenado em buffer, 
	// podemos enviar esses valores para o canal
	// sem um recebimento simultâneo correspondente.
	messages <- "armazenado"
	messages <- "canal"

	// Mais tarde, podemos receber esses dois valores
	// como de costume.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
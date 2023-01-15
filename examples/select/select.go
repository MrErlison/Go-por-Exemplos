// O _select_ do Go permite que você aguarde múltiplas
// operações de canal. Combinar goroutines e canais com
// _select_ é um recurso poderoso do Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Para o nosso exemplo, selecionaremos dois canais.
	c1 := make(chan string)
	c2 := make(chan string)

	// Cada canal receberá um valor após um período de
	// tempo, para simular, por exemplo, bloqueo de
	// operações RPC em execução em goroutines
	// simultâneas.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "um"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "dois"
	}()

	// Usaremos _select_ para aguardar esses dois valores
	// simultaneamente, imprimindo cada um assim que
	// chegar.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recebido", msg1)
		case msg2 := <-c2:
			fmt.Println("recebido", msg2)
		}
	}
}

// Os tempos limites (_timeouts_) são importantes para 
// programas que se conectam a recursos externos ou que,
// de outra forma, precisam limitar o tempo de execução.
// Implementar _timeouts_ no Go é fácil e elegante graças
// aos canais e ao `select`.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Para o nosso exemplo, suponha que estejamos executando
	// uma chamada externa que retorne seu resultado em um 
	// canal `c1` após 2s. Observe que o canal está armazenado
	// em buffer, portanto, o envio goroutine não está 
	// bloqueado. Este é um padrão comum para evitar vazamentos
	// de goroutine, caso o canal nunca seja lido.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Aqui está o `select` implementando um _timeout_.
	// `res := <-c1` aguarda o resultado e `<-time.After`
	// aguarda um valor a ser enviado após o tempo limite
	// de 1s. Como o `select` prossegue com o primeiro
	// recebimento que está pronto, assumiremos o caso
	// de tempo limite se a operação levar mais do que
	// 1s permitido.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Se permitirmos um _timeout_ maior de 3s, o 
	// recebimento de `c2` será bem-sucedido e o resultado
	// será impresso.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

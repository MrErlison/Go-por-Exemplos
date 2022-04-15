// Ao usar canais como parâmetros de função, você pode 
// especificar se um canal deve apenas enviar ou receber
// valores. Essa especificidade aumenta a segurança do
// tipo do programa.

package main

import "fmt"

// Esta função `ping` só aceita um canal para envio de
// valores. Seria um erro `compile-time` tentar receber
// neste canal.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// A função `pong` aceita um canal para recebimentos
// (`pings`) e um segundo para envios (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

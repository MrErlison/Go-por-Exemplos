// _Canais_ são os tubos que conectam goroutines 
// simultâneas. Você pode enviar valores pelos canais
// de uma goroutines e receber esses valores em outra
// goroutines.

package main

import "fmt"

func main() {

	// Crie um novo canal com `make(chan val-type).` 
	// Os canais são tipados pelos valores que transmitem.
	messages := make(chan string)

	// Envie um valor por um canal usando a sintaxe 
	// do canal `<-`. Aqui enviamos "ping" para o canal
	// de mensagens que fizemos acima, a partir de uma 
	// nova goroutines.
	go func() { messages <- "ping" }()

	// A sintaxe `<-` recebe um valor do canal. Aqui 
	// vamos receber a mensagem de "ping" que enviamos
	// acima e imprimi-la.
	msg := <-messages
	fmt.Println(msg)
}

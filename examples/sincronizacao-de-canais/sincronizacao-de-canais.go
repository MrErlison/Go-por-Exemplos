// Podemos usar canais para sincronizar a execução entre
// goroutines. Aqui está um exemplo de como usar um 
// recebimento de bloqueio para esperar que uma goroutine
// termine. Ao esperar que várias goroutines terminem, 
// você pode preferir usar um [WaitGroup](waitgroups).
package main

import (
	"fmt"
	"time"
)

// Esta é a função que executaremos em uma goroutine.
// O canal `done` será usado para notificar outra 
// goroutine de que o trabalho desta função está feito.
func worker(done chan bool) {
	fmt.Print("trabalhando...")
	time.Sleep(time.Second)
	fmt.Println("feito")

	// Envie um valor para notificar que terminamos.
	done <- true
}

func main() {

	// Inicie uma goroutine `worker`, dando a ele o 
	// canal para notificar.
	done := make(chan bool, 1)
	go worker(done)

	// Bloqueie até recebermos uma notificação do
	// trabalhador no canal.
	<-done
}

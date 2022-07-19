// A [limitação de taxa](http://en.wikipedia.org/wiki/Rate_limiting)
// é um mecanismo importante para controlar a utilização dos 
// recursos e manter a qualidade do serviço. O Go suporta
// a limitação de taxa com goroutines, [canais](canais) e 
// [tickers](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Primeiro, veremos a limitação básica da taxa. 
	// Suponha que queremos limitar o processamento de
	// solicitações recebidas. Atenderemos essas 
	// solicitações de um canal com o mesmo nome.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Este canal `limiter` receberá um valor a cada 
	// 200 milissegundos. Este é o regulador do nosso 
	// esquema de limitação de taxas.
	limiter := time.Tick(200 * time.Millisecond)

	// Ao bloquear um recebimento do canal `limiter` 
	// antes de atender a cada solicitação, nós 
	// limitamos a 1 solicitação a cada 200 milissegundos.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Podemos querer permitir pequenas rajadas de 
	// solicitações em nosso esquema de limitação de
	// taxas, preservando o limite geral. 
	// Podemos fazer isso protegendo o canal 
	// `limiter`. Este canal `burstyLimiter` 
	// permitirá rajadas de até 3 eventos.
	burstyLimiter := make(chan time.Time, 3)

	// Preencha o canal para representar o estouro permitido.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// A cada 200 milissegundos, tentaremos adicionar
	// um novo valor ao `burstyLimiter`, até o seu 
	// limite de 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Agora simule mais 5 solicitações recebidas. 
	// Os primeiros 3 deles se beneficiarão da capacidade
	// de rajadas do `burstyLimiter`.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

// Para esperar que várias goroutines terminem, podemos
// usar um grupo de espera (*wait group*).
package main

import (
	"fmt"
	"sync"
	"time"
)

// Esta é a função que executaremos em todas as goroutine.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Aqui o `Sleep` serve para simular uma tarefa
	// mais demorada.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// Esse `WaitGroup` é usado para esperar que todas as 
	// goroutines inicadas aqui terminem. 
	// Se um WaitGroup` for explicitamente passado para 
	// funções, isso deve ser feito usando ponteiro.
	var wg sync.WaitGroup

	// Inicie várias goroutines e aumente o contador
	// `WaitGroup` para cada uma.
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Evite a reutilização do mesmo valor `i` em
		// cada goroutine `closure`. Consulte o [FAQ](https://golang.org/doc/faq#closures_and_goroutines) 
		// para mais detalhes.
		i := i

		// Envolva a chamada do `worker` no `closure` e certifique
		// de dizer ao `WaitGroup` que o `worker` está pronto. 
		// Dessa forma, o próprio `worker` não precisa
		// saber das concorrência envolvidas em sua execução.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Bloqueie até que o contador `WaitGroup` volte para 0;
	// todos os `worker` notifiquem que terminaram.
	wg.Wait()

	// Observe que essa abordagem não tem uma maneira direta
	// de propagar os erros dos `worker`. Para casos de mais 
	// avançados, considere usar o [pacote errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup).
}

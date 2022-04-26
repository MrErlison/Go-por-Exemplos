// Neste exemplo, veremos como implementar uma fila
// de trabalhas usando goroutines e canais.
package main

import (
	"fmt"
	"time"
)

// Aqui está o `worker`, do qual executaremos várias 
// instâncias simultâneas. Esses worker receberão 
// trabalho no canal `jobs` e enviarão os resultados
// correspondentes em `results`. Aguardaremos um segundo
// por trabalho para simular uma tarefa dispendiosa.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// Para usar nosso grupo de `workers`, precisamos 
	// enviar o trabalho e coletar seus resultados. Criamos
	// 2 canais para isso.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Isso inicia 3 `workers`, inicialmente bloqueados
	// porque ainda não há nenhum `job`.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Aqui enviamos 5 `jobs` e depois fechamos esse 
	// canal para indicar que esse é todo o trabalho
	// que temos.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finalmente, coletamos todos os resultados do
	// trabalho. Isso também garante que as rotinas do
	// `worker` terminou. Uma maneira alternativa de
	// esperar por várias goroutines é usar um WaitGroup.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}


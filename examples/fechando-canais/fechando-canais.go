// Fechar um canal indica que não serão enviados mais
// valores nele. Isso pode ser útil para comunicar a
// conclusão aos receptores do canal.

package main

import "fmt"

// Neste exemplo, usaremos um canal `jobs` para comunicar
// o trabalho a ser feito da goroutine `main()` para uma
// goroutine `worker`. Quando não tivermos mais `jobs`
// para o `worker`, fecharemos o canal `jobs`.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Aqui está a goroutine do `worker`. Ela recebe
	// repetidamente `jobs` com `j, more := <-jobs`.
	// Nesta forma especial de recebimento de `2-value`,
	// `more` será falso se os `jobs` tiverem sido fechados
	// e todos os valores no canal já tiverem
	// sido recebidos. Usamos isso para notificar com `done`
	// em todos os nossos `jobs`.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Isso envia três trabalhos para o `worker` pelo
	// canal `jobs` e, em seguida, fecha.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Aguardamos o `worker` usando a abordagem de
	// sincronização que vimos anteriormente.
	<-done
}

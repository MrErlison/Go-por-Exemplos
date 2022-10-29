// O principal mecanismo para gerenciar estado no Go é
// a comunicação por canais. Vimos isso, por exemplo, 
// com [filas trabalhados](fila-de-trabalhos). 
// No entanto, existem algumas outras opções para gerencia de estado.
// Aqui vamos analisar o uso do pacote `sync/atomic` para 
// contadores atômicos acessados por múltiplas gorotinas.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Usaremos um número inteiro sem sinal para 
	// representar nosso contador (sempre positivo).
	var ops uint64

	// Um WaitGroup nos ajudará a esperar que todas as
	// gorotinas terminem seu trabalho.
	var wg sync.WaitGroup

	// Começaremos com 50 gorotinas que cada uma 
	// incrementa o contador exatamente 1000 vezes.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// Para incrementar atomicamente o contador,
				// usamos o `AddUint64`, passando o endereço
				// de memória do nosso contador `ops` 
				// com a sintaxe `&`.
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// Espere até que todas as gorotinas 
	// estejam prontas.
	wg.Wait()

	// É seguro acessar `ops` agora porque sabemos
	// que nenhuma outra gorotina está escrevendo 
	// nela. Também é possível ler com segurança 
	// enquanto elas são atualizadas, usando funções
	// como `atomic.LoadUint64`.
	fmt.Println("ops:", ops)
}

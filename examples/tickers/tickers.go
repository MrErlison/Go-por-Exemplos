// Os [temporizadores](temporizador) são para quando 
// você quiser fazer algo no futuro - _tickers_ são 
// para quando você quiser fazer algo repetidamente em
// intervalos regulares. Aqui está um exemplo de um 
// ticker que marca periodicamente até o pararmos.
package main

import (
	"fmt"
	"time"
)

func main() {

	// Os Tickers usam um mecanismo similar aos 
	// temporizadores: um canal que recebe valores. 
	// Aqui usaremos o `select` embutido no canal para
	// esperar os valores à medida que eles chegam a 
	// cada 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick em", t)
			}
		}
	}()

	// Tickers podem ser parados como temporizadores. 
	// Depois que um ticker for interrompido, ele não 
	// receberá mais valores em seu canal. Vamos parar 
	// o nosso depois de 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker parado")
}

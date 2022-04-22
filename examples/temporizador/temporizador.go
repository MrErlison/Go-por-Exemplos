// Muitas vezes queremos executar o código Go em algum
// momento no futuro, ou repetidamente em algum 
// intervalo. Os recursos integrados de temporizador 
// _timer_ e _ticker_ do Go facilitam essas duas 
// tarefas. Vamos olhar primeiro para os temporizadores
// e depois para os [tickers](tickers).
package main

import (
	"fmt"
	"time"
)

func main() {

	// Os temporizadores representam um único evento no 
	// futuro. Você diz ao cronômetro quanto tempo quer
	// esperar e ele fornece um canal que será notificado
	// naquele momento. Este cronômetro vai esperar 
	// 2 segundos.
	timer1 := time.NewTimer(2 * time.Second)

	// O `<-timer1.C` bloqueia o canal `C` do temporizador
	// até que ele envie um valor indicando que o 
	// temporizador disparou.
	<-timer1.C
	fmt.Println("Timer 1 disparou")

	// Se você quisesse apenas esperar, poderia ter usado
	// `time.Sleep`. Uma das razões pelas quais um
	// cronômetro pode ser útil é que você pode cancelar
	// o cronômetro antes que ele seja acionado. 
	// Aqui está um exemplo disso.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 disparou")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 parou")
	}

	// Dê ao `timer2` tempo suficiente para disparar, 
	// para mostrar que está de fato parado.
	time.Sleep(2 * time.Second)
}

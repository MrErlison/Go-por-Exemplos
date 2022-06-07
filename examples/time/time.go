// O Go oferece amplo suporte para tempos e durações;
// aqui estão alguns exemplos.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Começaremos obtendo a hora atual.
	now := time.Now()
	p(now)

	// Você pode criar uma estrutura de tempo fornecendo
	// o ano, mês, dia, etc. Os horários estão sempre 
	// associados a um local, ou seja, fuso horário.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Você pode extrair os vários componentes do valor
	// do tempo como esperado.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// O dia da semana, de segunda a domingo, também
	// está disponível.
	p(then.Weekday())

	// Esses métodos se comparam dois tempos, testando
	// se a primeira ocorre antes, depois ou ao mesmo
	// tempo que a segunda, respectivamente.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Os métodos `Sub` retornam uma duração que representa
	// o intervalo entre dois tempos.
	diff := now.Sub(then)
	p(diff)

	// Podemos calcular a tamanho da duração em várias
	// unidades
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Você pode usar `Add` para avançar um tempo por
	// uma determinada duração, ou com um `-` para
	// retroceder por uma duração.
	p(then.Add(diff))
	p(then.Add(-diff))
}

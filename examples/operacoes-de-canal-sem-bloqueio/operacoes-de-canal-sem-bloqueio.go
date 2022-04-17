// Os envios e recebimentos básicos nos canais estão
// bloqueados. No entanto, podemos usar `select` com uma
// cláusula `default` para implementar envios, recebimentos
// e até mesmo seleções multidirecionais sem bloqueio.
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Aqui está uma recepção sem bloqueio. Se um valor
	// estiver disponível nas `messages`, selecione
	// pegará o caso `<-messages` com esse valor.
	// Caso contrário, assumirá imediatamente o caso `default`.
	select {
	case msg := <-messages:
		fmt.Println("mensagem recebida", msg)
	default:
		fmt.Println("nenhuma mensagem recebida")
	}

	// Aqui, o `msg` não pode ser enviado para o canal
	// `messages`, porque o canal não tem buffer e não
	// há receptor. Portanto, o caso `default` está
	// selecionado.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("mensagem enviada", msg)
	default:
		fmt.Println("nenhuma mensagem enviada")
	}

	// Podemos usar vários `case` da cláusula `default`
	// para implementar uma seleção multidirecional
	// sem bloqueio. Aqui tentamos não bloquear
	// recebimentos em mensagens e sinais.
	select {
	case msg := <-messages:
		fmt.Println("mensagem recebida", msg)
	case sig := <-signals:
		fmt.Println("sinal recebido", sig)
	default:
		fmt.Println("nenhuma atividade")
	}
}

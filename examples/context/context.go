// No exemplo anterior, vimos como configurar um simples
// servidor HTTP. Os servidores HTTP são úteis para
// demonstrar o uso do context.Context para controlar
// a cancelamento. Um Context carrega prazos, sinais de
// cancelamento e outros valores de escopo de solicitação
// através de fronteiras de API e goroutines.

package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// Um `context.Context` é criado para cada solicitação
	// pela mecânica de `net/http` e está disponível com
	// o método `Context()`.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Aguarde alguns segundos antes de enviar uma resposta
	// para o cliente. Isso poderia simular algum trabalho
	// que o servidor está fazendo. Enquanto trabalha, fique
	// de olho no canal `Done()` do contexto por um sinal de
	// que devemos cancelar o trabalho e retornar o mais
	// rápido possível.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		// O método Err() do contexto retorna um erro que
		// explica porque o canal Done() foi fechado.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// Como antes, registramos nosso manipulador na rota "/hello"
	// e iniciamos o atendimento.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

// A biblioteca padrão do Go vem com um excelente 
// suporte para clientes e servidores HTTP no 
// pacote `net/http`. Nesse exemplo, vamos usá-lo
// para simples requisições HTTP.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Solicita HTTP GET para um servidor. `http.Get` é
	// um atalho conveniente para criar um objeto
	// `http.Client` e chamar seu método `Get`; ele usa 
	// o objeto `http.DefaultClient`, que possui 
	// padrões de configurações úteis.
	resp, err := http.Get("http://goporexemplos.github.io/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Exibe o status da resposta HTTP.
	fmt.Println("Response status:", resp.Status)

	// Exibe as primeiras 5 linhas do corpo da resposta.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

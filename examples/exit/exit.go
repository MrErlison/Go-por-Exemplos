// Use `os.Exit` para sair imediatamente com um status 
// determinado.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `defer` não será executado ao usar `os.Exit`, 
	// então esse `fmt.Println` nunca será chamado.
	defer fmt.Println("!")

	// Sair com status 3
	os.Exit(3)
}

// Observe que, em Go não se usa um valor de retorno 
// inteiro da função `main` para indicar o status de
// saída. Se você quiser sair com um status diferente
// de zero, você deve usar `os.Exit`.

// O `panic` normalmente significa que algo deu 
// inesperadamente errado. Principalmente, nós o usamos
// para falhar rapidamente em erros que não devem ocorrer
// durante a operação normal, ou que não estamos 
// preparados para lidar graciosamente.

package main

import "os"

func main() {

	// Usaremos `panic` em todo este site para verificar
	// se há erros inesperados. Este é o único programa 
	// no site projetado para entrar em pânico.
	panic("um problema")

	// Um uso comum do `panic` é abortar se uma função
	// retornar um valor de erro que não sabemos como 
	// (ou queremos) lidar. Aqui está um exemplo de 
	// `panic` se recebermos um erro inesperado ao criar
	// um novo arquivo.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

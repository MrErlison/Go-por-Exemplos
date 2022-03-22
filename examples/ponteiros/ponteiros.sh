# A `zeroval` não altera o `i` na função principal,
# `main`, mas a `zeroptr` muda porque tem uma referência
# ao endereço de memória dessa variável
$ go run ponteiros.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100

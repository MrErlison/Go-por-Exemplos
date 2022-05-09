# Se você executar `exit.go` usando `go run`, a 
# saída será retirada por `go` e exibida.
$ go run exit.go
exit status 3

# Ao criar e executar um binário, você pode ver o
# status no terminal.
$ go build exit.go
$ ./exit
$ echo $?
3

# Observe que o `!` do nosso programa nunca foi exibido.

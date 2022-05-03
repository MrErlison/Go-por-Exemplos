# Executar este programa fará com que ele entre em falha,
# imprima uma mensagem de erro e traces do goroutines,
# e saia com um status diferente de zero.

# Quando entra em falha pela primeira vez, o programa sai
# sem atingir o resto do código. Se você quiser ver o 
# programa tentar criar um arquivo temporário, comente o
# primeiro `panic`.
$ go run panic.go
panic: um problema

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# Observe que, ao contrário de algumas linguagens que 
# usam exceções para lidar com muitos erros, no Go é
# idiomático usar valores de retorno que indicam erros 
# sempre que possível.

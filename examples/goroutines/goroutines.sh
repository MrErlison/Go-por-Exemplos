# Quando executamos este programa, vemos a saída da 
# chamada de bloqueio primeiro, depois a saída das duas
# rotinas. A saída das goroutines pode ser intercalada,
# porque as goroutines estão sendo executadas 
# simultaneamente pelo tempo de execução Go.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# Em seguida, veremos um complemento para as rotinas
# em programas Go simultâneos: canais.
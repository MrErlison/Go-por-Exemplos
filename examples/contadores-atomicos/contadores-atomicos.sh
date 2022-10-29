# Esperamos obter exatamente 50.000 operações. Se usarmos 
# o não atômico `ops++` para incrementar o contador, 
# provavelmente teríamos um número diferente, mudando entre 
# execuções, porque as gorotinas interfeririam umas com as 
# outras. Além disso, obteriamos falhas de dados na execução
# com a flag `-race`.
$ go run contadores-atomicos.go
ops: 50000

# Em seguida, analisaremos os mutexes, outra ferramenta
# para gerenciar estado.
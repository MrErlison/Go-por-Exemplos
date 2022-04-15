# Recebemos os valores `"um"` e depois `"dois"` 
# como esperado
$ time go run select.go 
recebido um
recebido dois

# Observe que o tempo total de execução é de apenas ~2
# segundos, já que os `Sleeps` de 1 e 2 segundos são 
# executados simultaneamente.
real	0m2.245s

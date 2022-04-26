# Nosso programa em execução mostra os 5 trabalhos que
# estão sendo executados. O programa leva apenas cerca de
# 2 segundos, apesar de fazer cerca de 5 segundos de
# trabalho total, porque há 3 trabalhos operando 
# simultaneamente.

$ time go run fila-de-trabalhos.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s

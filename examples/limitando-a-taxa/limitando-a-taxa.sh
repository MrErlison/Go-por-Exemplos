# Executando nosso programa, vemos o primeiro lote
# de solicitações tratadas uma vez a cada ~200 
# milissegundos, conforme desejado.
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

# Para o segundo lote de solicitações, atendemos os
# 3 primeiros imediatamente por causa da limitação
# da taxa explosiva e, em seguida, atendemos os 2
# restantes com atrasos de ~200 ms cada.
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC

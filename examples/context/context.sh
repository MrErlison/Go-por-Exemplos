# Execute o servidor em segundo plano.
$ go run context-in-http-servers.go &

# Simule uma solicitação de cliente para 
# /hello, pressionando Ctrl+C logo após 
# iniciar para sinalizar cancelamento.
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended

# Quando executamos o programa, a mensagem "ping" é 
# passada com sucesso de uma goroutine para outra 
# através do nosso canal.
$ go run canais.go 
ping

# Por padrão, envia e recebe blocos até que o 
# remetente e o destinatário estejam prontos. Essa 
# propriedade nos permitiu esperar no final do nosso 
# programa pela mensagem "ping" sem ter que usar qualquer
# outra sincronização.
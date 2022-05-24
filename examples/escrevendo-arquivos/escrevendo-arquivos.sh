# Tente executar o código de gravação de arquivos.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# Em seguida, verifique o conteúdo dos arquivos.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# Em seguida, analisaremos a aplicação de algumas
# das ideias de E/S de arquivos que acabamos de ver nos
# fluxos `stdin` e `stdout`.
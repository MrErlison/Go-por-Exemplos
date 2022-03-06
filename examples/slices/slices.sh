# Observe que, embora os `slices` sejam de tipos 
# diferentes dos `arrays`, eles são tratados de forma
# semelhante pelo fmt.Println.
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# Confira esta ótima postagem do [blog da equipe Go](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html) 
# para obter mais detalhes sobre o design e 
# implementação do `slice` no Go

# Agora que vimos `array` e `slices`, veremos a outra
# estrutura de dados integrada do Go: mapas.
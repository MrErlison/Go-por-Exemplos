# Go por Exemplos

Aprenda a programar por meio de exemplos em GO com anotações. Este projeto é um fork do projeto original [Go by Example](https://github.com/mmcgrana/gobyexample). 

# Panorama geral

O site [Go por Exemplo](https://goporexemplos.github.io/) é construído extraindo código e comentários de arquivos exemplos e renderizando-os através dos `templates` no diretório estático `public`. Os programas que implementam esse processo de compilação estão na pasta `tools`, juntamente com dependências especificadas no arquivo `go.mod`.

O diretório `public` construído pode ser disponibilizado por qualquer sistema de conteúdo estático. O local de produção usa S3 e CloudFront, por exemplo.

### Construindo

[![Build Status](https://github.com/MrErlison/Go-por-Exemplos/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/MrErlison/Go-por-Exemplos/actions)

Para construir seu site você vai precisar do Go instalado. Execute:

```console
$ tools/build
```

Para contrução contínua:

```console
$ tools/build-loop
```

Para visualizar o site localmente:

```
$ tools/serve
```

depois acesse o endereço `http://127.0.0.1:8000/` no seu navegador.

### Publicando

Para publicar o site:

```console
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ tools/upload
```

### Licença

Este trabalho protegido por direitos autorais por Mark McGranaghan e está licenciado pela
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

O Go Gopher é protegido por direitos autorais de [Renée French](http://reneefrench.blogspot.com/) e licenciado pela 
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Traduções

Traduções de colaboradores do site Go by Example estão disponíveis em:


* [Chinese](https://gobyexample-cn.github.io/) por [gobyexample-cn](https://github.com/gobyexample-cn)
* [Czech](http://gobyexamples.sweb.cz/) por [martinkunc](https://github.com/martinkunc/gobyexample-cz)
* [French](http://le-go-par-l-exemple.keiruaprod.fr) por [keirua](https://github.com/keirua/gobyexample)
* [Italian](http://gobyexample.it) pela [Go Italian community](https://github.com/golangit/gobyexample-it)
* [Japanese](http://spinute.org/go-by-example) por [spinute](https://github.com/spinute)
* [Portuguese](https://goporexemplos.github.io) por [Mr. Erlison](https://github.com/MrErlison/goporexemplos/)
* [Korean](https://mingrammer.com/gobyexample/) por [mingrammer](https://github.com/mingrammer)
* [Russian](https://gobyexample.com.ru/) por [badkaktus](https://github.com/badkaktus)
* [Spanish](http://goconejemplos.com) pela [Go Mexico community](https://github.com/dabit/gobyexample)
* [Ukrainian](http://butuzov.github.io/gobyexample/) by [butuzov](https://github.com/butuzov/gobyexample)

### Agradecimentos

Obrigado ao [Jeremy Ashkenas](https://github.com/jashkenas)
para [Docco](http://jashkenas.github.io/docco/), pela inspição neste projeto.


[FAQ](https://github.com/mmcgrana/gobyexample#faq)

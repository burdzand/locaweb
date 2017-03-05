# Locaweb Twitter

A aplicação foi escrita em GO e se encontra na Pasta /go/src/github.com/andersondborges/locaweb

Caso queirar rodar em máquina local deve-se ter instalado * [Go](https://golang.org/doc/install)
Para instalar vá na pasta do projeto e  basta executar o camando "go build"
e em seguida executar o arquivo ./locaweb

Caso deseje que seja executado pelo PATH basta ir na pasta do projeto executar o comando "go install"
e em seguida executar commando "locaweb ou /go/bin/locaweb"

Para executar os testes  vá no dirtório do projeto e execute "go test -v"

O acessado através de `http://localhost:9090`

#Instação Docker
O projeto foi desenvolvido sem nenhuma biblioteca de terceiros, logo não haverá necessidade de instalar pacotes dependentes.

Dentro do diretório raiz do projeto, que é onde está o arquivo `Dockerfile`, 
execute o comando `docker build -t andersondborges/locaweb:1.0` para que o container seja criado.

Para levantar o Conteiner e executar a aplicação use o commando `docker run -p 9090:9090 -d --name locaweb andersondborges/locaweb:1.0
o serviço estará desponível na porta 9090.



# O início

## Pacotes
Go é baseado em pacotes.

O Pacote *main* é o pacote que diz ao compilador que o programa em questão é um executável. Ele é o ponto de entrada de uma aplicação.

```golang
package main
```

## Workspace

O Worskpace do Go é definido pela variável de ambiente ```GOPATH```

Todos seus projetos e códigos Go devem ser escritos na pasta definida pelo ```GOPATH``` setado.

O Go procura por pacotes a partir da raiz do ```GOPATH``` e do ```GOROOT```.

```GOROOT``` é o local onde foi feita a  instalação do Go.

Uma boa prática é criar uma pasta com o nome ```go``` na raiz da pasta de seu usuário no sistema operacional.

Exemplo de ```GOPATH``` e ```GOROOT``` no linux
```
GOPATH="/home/diego/go"
GOROOT="/usr/local/go"
```

# Hello World

Crie um arquivo chamado *main.go* com o seguinte conteúdo:

```golang
package main

import (
 "fmt"
)

func main(){
  fmt.Println("Hello World!")
}
```

No exemplo acima, a primeira linha indica que este arquivo Go pertence ao pacote *main*.

Logo após temos uma declaração de importação de outro pacote, o *fmt*. O pacote *fmt* prove funções de formatação de I/O.

Importamos o pacote utilizando a palavra reservada *import*.

```func main()``` é o ponto principal de entrada para aplicação.

```Println``` é a função do pacote *fmt* que imprime "hello world".

Vamos compilar o código:
```
go build main.go
```

Isto vai gerar um binário chamado *main*.
Vamos executar ele:
```
./main
```

Existe uma maneira mais simples de rodar o código go, que abstrai a compilação.

```
go run main.go
```
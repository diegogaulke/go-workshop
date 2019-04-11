# O início

## Pacotes
Go é baseado em pacotes.

O Pacote *main* é o pacote que diz ao compilador que o programa em questão é um executável. Ele é o ponto de entrada de uma aplicação.

```
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

```
package main

import (
 "fmt"
)

func main(){
  fmt.Println("Hello World!")
}
```


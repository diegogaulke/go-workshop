# Concorrência

Go foi criado pensando em concorrência.

A concorrência em Go pode ser alcançada utilizando as *go routines*, que são *threads* leves.

### Go routine

Go routines são funções que podem rodar em paralelo, ou concorrentemente com outra função.
Criar uma Go routine é muito simples.

Simplesmente adicionamos a palavra reservada *go* na frente de uma chamada de função, assim fazemos ela ser executada em paralelo.
Go routines são muito leves, por isso podemos criar milhares dela.

```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	go c()
	fmt.Println("I am main")
	time.Sleep(time.Second * 2)
}
func c() {
	time.Sleep(time.Second * 2)
	fmt.Println("I am concurrent")
}

//=> I am main
//=> I am concurrent
```

Como podemos ver no exemplo acima, a função *c* é uma Go routine que executa em paralelo com a função main.
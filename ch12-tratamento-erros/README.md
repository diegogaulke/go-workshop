# Tratamento de Erros

Erros são resultados inesperados (`runtime`) durante a execução de programa. Uma chamada de API pode resultar em **sucesso** ou **falha**.

Um erro em Go pode ser reconhecido quando um *tipo* `error` aparece.

```golang
resp, err := http.Get("http://example.com/")
```

No código acima, o resultado da chamada pode resultar em um erro. Imagine se o dominio *example.com* não exista. Um erro será retornado.

No código abaixo fazemos uma checagem para verificar se ocorreu erro na chamada e tratar o mesmo.

```golang
package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
```

### Retornando um erro customizado em uma função

Quando estamos escrevendo uma função própria, podem existir casos que precisamos lançar erros.
Esses erros podem ser retornados com a ajuda do objeto `error`.

```golang
package main

import (
	"errors"
	"fmt"
)

func Increment(n int) (int, error) {
	if n < 0 {
		// return error object
		return 0, errors.New("math: cannot process negative number")
	}
	return (n + 1), nil
}
func main() {
	num := 5

	if inc, err := Increment(num); err != nil {
		fmt.Printf("Failed Number: %v, error message: %v", num, err)
	} else {
		fmt.Printf("Incremented Number: %v", inc)
	}
}
```

A maioria dos pacotes nativos do Go, ou pacotes externos que usamos, tem algum mecanismo de tratamento de erros. Em qualquer função que chamamos podemos nos deparar com o retorno de um erro, que precisa ser tratado ou ignorado.

### Panic

A função `panic()` interrompe a execução do programa. Ela aceita como parâmetro um tipo `error` ou uma `string`. Ela só deve ser chamada quando um problema totalmente inesperado que impede o prosseguimento da execução da aplicação é encontrado.

### Defer

O `defer` é algo que sempre vai ser executado após o término de uma função.

```golang
//Go
package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
```

No exemplo acima, chamamos a função `panic()`. Como pode ser visto, há uma declaração de `defer` que fará com que o programa execute a linha no final da execução do programa. O `defer` também pode ser usado quando precisamos que algo seja executado ao final de uma função, por exemplo, fechar um arquivo.

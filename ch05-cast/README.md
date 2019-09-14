# Cast de Tipos

Um tipo de dado pode ser convertido em outro usando *type casting*.

Vamos ver um exemplo:

```golang
package main

import "fmt"

func main() {
    a:= 1.1 // float
    fmt.Println(a) // imprime 1.1

    b:= int(a)
    fmt.Println(b) // imprime 1
}
```

Convertendo numéricos para string e vice-versa:

```golang
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	a := 1         // int
	fmt.Println(a) // imprime 1

	s := strconv.Itoa(a)
	fmt.Println(reflect.TypeOf(s)) // imprime o tipo da variavel string
}
```

Praticamente todos tipos de dados podem ser convertidos. Você deve garantir que o tipo de dado possa ser convertido para o tipo desejado.

Observe que no exemplo acima usamos o pacote nativo *strconv* para efetuarmos a conversão de *string*.
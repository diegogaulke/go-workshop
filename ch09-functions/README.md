# Funções

A função `main()` definida no pacote de mesmo nome, é o ponto de entrada para a execução de um programa Go.

Podemos definir mais funções e as utilizar, conforme este exemplo:

```golang
package main

import "fmt"

func add(a int, b int) int {
	c := a + b
	return c
}
func main() {
	fmt.Println(add(2, 1)) // imprime 3
}
```

Como podemos ver no exemplo anterior, uma função é definida pela palavra `func` seguida pelo nome da função. Os argumentos que uma função precisa, devem ser definidos de acordo com seu tipo de dado, e em seguida o tipo de retorno.

Também podemos definir o retorno na declaração da função:

```golang
package main

import "fmt"

func add(a int, b int) (c int) {
	c = a + b
	return
}
func main() {
	fmt.Println(add(2, 1)) // imprime 3
}
```

No exemplo acima, a variável `c` é definida na declaração da função, não sendo necessária ser informada na declaração de retorno ao final.

Você também pode definir múltiplas variáveis de retorno para uma função:

```golang
package main

import "fmt"

func add(a int, b int) (int, string) {
	c := a + b
	return c, "successfully added"
}
func main() {
	sum, message := add(2, 1)
	fmt.Println(message)
	fmt.Println(sum)
}
```
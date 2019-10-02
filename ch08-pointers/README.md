# Ponteiros

Golang provê a utilização de ponteiros.

Resumidamente, um ponteiro armazena um endereço de memória de um valor.

Um ponteiro é definido por `*`.

Um ponteiro é definido de acordo com seu tipo de dado, exemplo:

```golang
var ap *int
```

No código acima a variável `ap` é um ponteiro para um valor do tipo `int`.

O operador `&` pode ser utilizado para buscar o endereço de uma variável.

```golang
a := 12
ap := &a
```

O valor referenciado ao ponteiro pode ser acessado usando o operador `*`.

```golang
a := 12
ap := &a
fmt.Println(*ap) // imprime 12
```

Para ficar mais claro, no exemplo abaixo imprimimos o valor de referência do ponteiro, bem como seu endereço de memória.

```golang
package main

import "fmt"

func main() {
	a := 12
	ap := &a
	fmt.Println(*ap) // Imprime 12
	fmt.Println(ap) // Imprime um endereco de memoria, algo como 0xc0000180f0
}
```

Um exemplo mais completo:

```golang
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```
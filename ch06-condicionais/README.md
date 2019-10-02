# Condicionais

## if else

Em go, para declaração de condições *if* e *else* utilizamos a sintaxe muito semelhante ao Java e javascript, conforme o exemplo abaixo.

```golang
package main

import "fmt"

func main() {
    if num := 9; num < 0  {
        fmt.Println(num, "é negativo")
    } else if num < 10 {
        fmt.Println(num, "tem um dígito")
    } else {
        fmt.Println(num, "tem vários dígitos")
    }
}
```

## switch case

*switch* e *case* ajudam a organizar declarações condicionais.

No Go fazemos assim:

```golang
package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	default:
		fmt.Println("nenhum")
	}
}
```

> Obs.: No Go não precisa do `break`, ele é implicito.

Simples certo? :)
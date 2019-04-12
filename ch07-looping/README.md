# Looping

Go tem apenas uma *keyword* para iteração. É o *for*.

É isto mesmo, não temos *while*. A ideia é simplificar correto? 

Então vai um loop simplificado:

```golang
package main

import "fmt"

func main() {
	i := 0
	sum := 0
	for i < 10 {
		sum += 1
		i++
	}
	fmt.Println(sum)
}
```

Percebam que no loop acima não estamos incrementando o índice na declaração do for, então aí vai um loop "completo":

```golang
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)
}
```

Muito mais simples não?

Tudo bem, sabemos que você quer fazer um loop infinito, então aí vai:

```golang
package main

func main() {
	for {
	}
}
```

Uma iteração simples com um loop em array:

```golang
package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
```

## Looping com range

Existe uma forma de iteração sobre mapas e slices chamada de *range*.

Quando iteramos em uma slice, dois valores são retornados em cada iteração. O primeiro é o índice, o segundo é o valor do elemendo naquele índice.

```golang
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("valor em %d = %d\n", i, v)
	}
}
```

Você pode ignorar o índice ou o valor os apontando para ```_```.

```golang
for i, _ := range pow
for _, value := range pow
```

Se você quiser apenas usar o índice:

```golang
for i := range pow
```
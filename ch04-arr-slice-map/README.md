# Arrays, Slices e Maps

## Arrays

Array é uma sequencia de elementos do mesmo tipo de dados. Um array tem um tamanho fixo que é definido em sua declaração, e não pode ser mais alterado depois de declarado.


Elas podem ser declaradas assim:
```golang
var a [10]int
```

Arrays também podem ser multidimensionais.
```golang
var arr [3][3]int
```

Iniciando uma array com valores.
```golang
var arr [3]int{1, 7, 11}
```

Setando valores a uma array já definida.
```golang
var arr [3]int
arr[0] = 1
arr[2] = 9
```


Você pode usar ```...``` na definição de capacidade de uma array, e deixar o compilador definir sua capacidade baseada nos elementos na declaração.

```golang
var arr [...] int{1, 2}
```

No caso acima o tamanho da array vai ser de 2 elementos.

### Tamanho de uma array:

O tamanho de uma array pode ser encontrado usando a função nativa ```len()```.

```golang
package main

import(
    "fmt"
)

func main() {
    var arr [...]int{9, 0, 3}
    fmt.Println(len(arr)) // imprime 3
}
```

Visto tudo isto, podemos dizer:

Arrays são limitadas para casos específicos onde os valores podem mudar em tempo de execução.

Mas não se preocupe, para isto temos os *Slices*.

## Slices

Slices é *wrap* flexível e robusto que abstrai uma array. Em resumo, um slice não detém nenhum dado nele. Ele apenas referência arrays existentes.

A declaração de uma slice é parecida com a de uma array, mas sem a capacidade definida.

```golang
package main

import "fmt"

func main() {
    // declaracao com var
    var s []int
    fmt.Println(s)

    // shorthand
    s2 := []int
    fmt.Println(s2)

    // len de uma slice
    fmt.Println(len(s2))
}
```

O código acima criou uma slice sem capacidade inicial e sem nenhum elemento.

Criando uma slice a partir de uma array:

```golang
package main

import "fmt"

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    // cria uma slice de a[1] até a[3]
    var b []int = a[1:4]
    fmt.Println(b)
}
```

A sintáxe ```a[start:end]``` cria uma slice a partir da array ```a``` iniciando do índice ```start``` até o índice ```end - 1```.
Então na linha 8 do código acima, ```a[1:4]``` cria uma representação da array ```a``` iniciando do índice 1 até o 3. Sendo assim a slice ```b``` tem os valores [77 78 79].


Criando uma slice usando ```make()```, uma função nativa do Go, que cria uma array e retorna um slice referenciando a mesma.

```func make([]T, len, cap) []T``` pode ser usado para criar uma slice, passando como parâmetro o tipo, o tamanho e a capacidade. A capacidade é opcional, e caso não informada seu default é o ```len``` que é um campo obrigatório.

```golang
package main

import (  
    "fmt"
)

func main() {  
    i := make([]int, 5, 5)
    fmt.Println(i)
}
```
 
#### Adicionando elementos a uma slice

Como sabemos arrays são limitadas em seu tamanho e não podem ser aumentadas. Slices tem seu tamanho dinâmico e podem receber novos elementos em tempo de execução usando a função nativa ```append```.

A definição da função append é a seguinte
```func append(s []T, x ...T) []T```

A sintáxe, ```x ...T ``` significa que a função aceita um número variável de elementos no parâmetro x, desde que respeitem o tipo do slice.

Uma questão que pode ter *ficado no ar*: Se uma slice é um *wrap* de uma array, como ela tem esta flexibilidade?
Bem, isso acontecer *debaixo dos panos*. Por exemplo: quando um novo elemento é adicionado a um slice ocorre o seguinte:

1. Uma nova array é criada
2. Os elementos da array atual são copiados
3. O elemenado *adicionado* ao slice é incluido na array
4. É retornado um slice, que é uma referência a esta nova array

Pretty cool right? :)


## Maps

Em Go, *Maps* são o tipo de dado que mapeia uma chave a um valor, como em Java e Python.

<img src="https://m.media-amazon.com/images/M/MV5BNDU1NTMwNTEtNjk0Yy00NDNlLWFiODctZWI5ODVlMGZmNzk2XkEyXkFqcGdeQXVyNjcwMzEzMTU@._V1_.jpg" alt="Obvious" width="50%"/>

*Fucking obvious*

Declarando um *map* em Go

```golang
var m map[string]int
```

Acima, a variável *m* referência o mapa criado, no qual a chave é uma *string* e o valor é um *int*.

Abaixo alguns exemplos de uso de map.

```golang
package main

import "fmt"

func main() {
	m := make(map[string]bool)

	m["www.google.com.br"] = true
	m["www.facebook.com"] = true

	if m["www.orkut.com"] == false {
		fmt.Println("Oops, nunca acessamos este site")
    }
    
    m["www.orkut.com.br"] = true

    fmt.Println(len(m))

    delete(m, "www.orkut.com.br")

    fmt.Println(len(m))
}
```
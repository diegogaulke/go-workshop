# Method, Structs, and Interfaces

Go é orientado a objetos? Não.

Porém as *structs*, *interfaces* e *methods* fazem quem conhece orientação a objetos se sentir _em casa_.

## Struct

A struct é um agrupamento ou uma coleção, tipada, de diferentes campos. Ela é usada para agrupar dados. Por exemplo, se quisermos agrupar dados de um tipo *Usuario* podemos definir uma *struct* com id, login e senha.

```golang
type Usuario struct {
	Id    int
	Login string
	Senha string
}
```

Agora alguns exemplos de como criar um dado com a *struct* criada.

```golang
// modo 1: passando valores na sequencia que os atributos estao na struct
u1 := Usuario{1, "diego", "diegobgaulke@gopher.go"}

// modo 2: definindo atributos
u2 := Usuario{Id: 1, Login: "diego"}

// modo 3:
u3 := new(Usuario)
```

Podemos acessar os dados da struct utilizando ponto (.)

```golang
u1 := Usuario{1, "diego", "diegobgaulke@gopher.go"}
fmt.Println(u1.Login)
```

Você também pode acessar os atributos de uma *struct* utilizando diretamente ponteiro.

```golang
up = &Usuario{1, "diego", "diegobgaulke@gopher.go"}
fmt.Println(up.Login)
```

## Methods

*Methods*  são tipos especiais de função com um *receiver*. Um *receiver* pode ser tanto um valor como um ponteiro.

Vamos criar um *Method* que tem um *receiver* do tipo *Usuario* que criamos anteriormente.

```golang
package main

import "fmt"

// Usuario defination
type Usuario struct {
	Id    int
	Login string
	Senha string
}

// method defination
func (u *Usuario) descricao() {
	fmt.Printf("Id de %v é %v.", u.Login, u.Id)
}
func (u *Usuario) setLogin(l string) {
	u.Login = l
}

func (u Usuario) setSenha(s string) {
	u.Senha = s
}

func main() {
	up := &Usuario{Id: 1, Login: "diego", Senha: "xyz"}
	up.descricao()
	// => Id de diego é 1
	up.setLogin("dg")
	fmt.Println(up.Login)
	//=> dg
	up.setSenha("zyz")
	fmt.Println(up.Senha)
	//=> zyx
}
```

Como vimos no exemplo acima, agora o método *descricao()* pode ser chamado usando o operador ponto (*.*)
Repare que o *receiver* é um ponteiro. Com o ponteiro estamos passando uma referencia para o valor, então se fizermos alguma alteração no método ela irá refletir no *receiver* *up*.

Observe que no exemplo o valor da senha não é alterado, porque o método *setSenha()* não é do tipo do *receiver*, enquanto que o método *setLogin()* é um ponteiro.

## Interfaces


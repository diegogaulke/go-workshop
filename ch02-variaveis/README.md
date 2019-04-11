# Variáveis

Variáveis em Go são tipadas estáticamente, ou seja seu tipo é checado na hora de sua declaração, sendo imutável.
Elas são declaradas de maneira explícita.

```golang
var i int
```

No exemplo acima o valor inicial será zero. Para declarar um valor inicial para a variável você pode usar a sintáxe:

```golang
var i int = 1
```

Dentro de uma função o *short assignment* ```:=``` pode ser usado para uma declaração implícita de variável, sem a necessidade da palavra reservada ```var```.

```golang
i := 1
```

Também é possível declarar mais de uma variável na mesma linha se forem do mesmo tipo.

```golang
var x, y int = 800, 600
```
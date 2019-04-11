# Tipos de dados básicos

Como em todas as linguagens, Go suporta dezenas de tipos de dados.

Vamos explorar alguns deles.

Sequencia de caracteres: ```string```

Tipos numéricos inteiros: ```int8, uint8 (byte), int16, uint16, int32 (rune), uint32, int64, uint64, int, uint, uintptr```

Tipos numéricos de ponto flutuante: ```float32, float64```

Tipos complexos: ```complex64 complex128```

Observe que ```byte``` é um alias de ```uint8```, e ```rune``` é um alias de ```int32```.


Alguns exemplos de declaração:

```golang
var b bool = true
var i int = 1
var x int32 = 35000
var s string "hello world"
var f float32 3.14
```
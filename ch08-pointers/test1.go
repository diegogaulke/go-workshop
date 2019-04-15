package main

import "fmt"

func main() {
	a := 12

	altera(a)
	fmt.Println(a)

	alteraPointer(&a)
	fmt.Println(a)
}

func alteraPointer(i *int) {
	*i = 20
}

func altera(i int) {
	i = 20
}

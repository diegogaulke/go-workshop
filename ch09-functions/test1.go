package main

import "fmt"

func add(a int, b int) (c int, d int) {
	c = a + b
	return
}

func main() {
	i, d := add(2, 1)
	fmt.Println(i) // imprime 3
	fmt.Println(d) // imprime 0

	sum(1, 2, 3)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func test(a ...interface{}) {

}

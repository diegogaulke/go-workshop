package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 2
	}
	fmt.Println(sum)

	x := 0
	for {
		x++

		if x == 5 {
			continue
		}

		fmt.Println(x)

		if x > 10 {
			break
		}
	}
}

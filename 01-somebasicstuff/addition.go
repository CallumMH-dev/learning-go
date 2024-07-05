package main

import "fmt"

func main() {

	x := 10
	y := 10
	var equals int

	fmt.Println("hello")
	addition(&x, &y, &equals)
	print(equals)
}

func addition(sum1 *int, sum2 *int, equals *int) {
	*equals = *sum1 + *sum2
}

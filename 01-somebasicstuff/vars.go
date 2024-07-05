package main

import "fmt"

func main() {
	var x int = 5
	y := 10
	var name string = "Name"

	fmt.Println(x, y, name)
	storeAndChangeValue(&x)

}

func storeAndChangeValue(holder *int) {
	fmt.Println(holder)
	var temp int
	temp = *holder
	*holder = 200
	fmt.Println(*holder)
	fmt.Println(temp)
}

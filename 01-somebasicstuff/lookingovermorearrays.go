package main

import (
	"fmt"
)

func main() {

	fmt.Println("Looking over arrays, more specifically make, copy, append, slice!")

	arrayNum := [4]int{0, 3, 2, 1}
	fmt.Println(arrayNum)

	arrayString := [5]string{

		"ZERO",
		"ONE",
		"TWO",
		"THREE",
		"FOUR",
	}

	fmt.Println(arrayString)

	a := [1]string{"Hello"}
	b := [1]string{" world"}

	fmt.Println(a[0] + b[0])

	dynamicArr := []int{1, 2, 3, 4, 5}

	arrayCopy := make([]int, 5, 10)

	fmt.Println(arrayCopy)

	copy(arrayCopy, dynamicArr)

	fmt.Println(dynamicArr)
	fmt.Println(arrayCopy)

	arrayCopy[2] = 99

	fmt.Println(dynamicArr)
	fmt.Println(arrayCopy)
	slicedArr := arrayCopy[0:2]

	fmt.Println(slicedArr)

	appending := append(arrayCopy, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	//would have assumed to get an error since capacity of arrayCopy is 10 but maybe that isn't how that works

	// c := [5]int{1,2,3,4,5}
	// appendtest := append(c,1,2,3,4,5,6)
	// ah it must be a slice.

	c := make([]int, 0)
	appendtest := append(c, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	fmt.Println(appending)
	fmt.Println(len(appendtest))
	fmt.Println(appendtest)

}

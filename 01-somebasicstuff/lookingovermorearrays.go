// the notes in this are just for personal use and rambling and figuring out how things work, this is by no means a tutorial

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
	//can add arrays

	dynamicArr := []int{1, 2, 3, 4, 5}
	//unsure if you would call this a dynamic array or not but it feels like one since i'm not manually setting a limit

	arrayCopy := make([]int, 5, 10)

	fmt.Println(arrayCopy)

	copy(arrayCopy, dynamicArr)
	// talk more about copying down below

	fmt.Println(dynamicArr)
	fmt.Println(arrayCopy)

	arrayCopy[2] = 99

	//creating copies do not affect the original array

	fmt.Println(dynamicArr)
	fmt.Println(arrayCopy)

	slicedArr := arrayCopy[0:2]
	//slices element 0,1

	fmt.Println(slicedArr)

	appending := append(arrayCopy, 6, 7, 8)
	// c := [5]int{1,2,3,4,5}
	// appendtest := append(c,1,2,3,4,5,6)
	// ah it must be a slice.

	c := make([]int, 0)
	appendtest := append(c, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	// I think underlying arrays and make / append works like this;
	// var := make([]int,5) creates an object for an array
	// var1 :=append(objectvar,ints,ints,ints)
	// which creates var1 into an array of sorts??

	fmt.Println(appending)
	fmt.Println(len(appendtest))
	fmt.Println(appendtest)

	dynamicArr1 := []int{1, 2, 3, 4, 5}

	arrayCopy1 := make([]int, 2)

	copy(arrayCopy1, dynamicArr1)
	fmt.Println(arrayCopy1)
	fmt.Println(dynamicArr1)

	//while i don't know if i'll ever use the capacity section of make(int[],_,_) understanding what the first value does helps alot

}

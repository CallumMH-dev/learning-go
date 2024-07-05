package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var output int

	readLine := bufio.NewReader(os.Stdin)
	fmt.Println("since i don't know how to do switch statements at all i will sub with if statements")

	fmt.Println("enter first number")
	sum1, _ := readLine.ReadString('\n')

	fmt.Println("1 add, 2 sub, 3 multi, 4 divi")

	operator, _ := readLine.ReadString('\n')

	fmt.Println("enter second number")
	sum2, _ := readLine.ReadString('\n')

	println("this is sum1", sum1, "this is sum2", sum2, "this is the op", operator)

	//calculate(&sum1, &sum2, &operator, &output)

	fmt.Println(output)
}

func calculate(sum1 *int, sum2 *int, operator *int, output *int) {
	fmt.Println("before ifs")
	fmt.Println(*operator)
	if *operator == 1 {
		fmt.Println("inside 1")
		*output = *sum1 + *sum2
	} else {
		if *operator == 2 {
			fmt.Println("inside 2")
			*output = *sum1 - *sum2
		} else {
			if *operator == 3 {
				fmt.Println("inside 3")
				*output = *sum1 * *sum2
			} else {
				if *operator == 4 {
					fmt.Println("inside 4")
					*output = *sum1 / *sum2
				}
			}
		}
	}

}

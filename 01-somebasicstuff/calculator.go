package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var output int

	readLine := bufio.NewReader(os.Stdin)
	fmt.Println("since i don't know how to do switch statements at all i will sub with if statements")

	fmt.Println("enter first number")
	num1, _ := readLine.ReadString('\n')
	sum1, _ := strconv.Atoi(num1)

	fmt.Println("1 add, 2 sub, 3 multi, 4 divi")

	placeholder, _ := readLine.ReadString('\n')
	operator, _ := strconv.Atoi(placeholder)

	fmt.Println("enter second number")
	num2, _ := readLine.ReadString('\n')
	sum2, _ := strconv.Atoi(num2)

	calculate(&sum1, &sum2, &operator, &output)

	fmt.Println(output)
}

func calculate(sum1 *int, sum2 *int, operator *int, output *int) {

	if *operator == 1 {
		*output = *sum1 + *sum2
	} else {
		if *operator == 2 {
			*output = *sum1 - *sum2
		} else {
			if *operator == 3 {
				*output = *sum1 * *sum2
			} else {
				if *operator == 4 {
					*output = *sum1 / *sum2
				}
			}
		}
	}

}

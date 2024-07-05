package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var output int

	readLine := bufio.NewReader(os.Stdin)
	fmt.Println("since i don't know how to do switch statements at all i will sub with if statements")

	fmt.Println("enter first number")
	sum1, _ := readLine.ReadString('\n')
	trimUserData1 := strings.TrimRight(sum1, "\n")

	fmt.Println("1 add, 2 sub, 3 multi, 4 divi")
	operatorCapture, _ := readLine.ReadString('\n')
	trimOperator := strings.TrimRight(operatorCapture, "\n")

	fmt.Println("enter second number")
	sum2, _ := readLine.ReadString('\n')
	trimUserData2 := strings.TrimRight(sum2, "\n")

	num1, _ := strconv.Atoi(trimUserData1)
	num2, _ := strconv.Atoi(trimUserData2)
	operator, _ := strconv.Atoi(trimOperator)

	calculate(&num1, &num2, &operator, &output)

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

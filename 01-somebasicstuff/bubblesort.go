package main

import "fmt"

func main() {

	testData := [10]int{1, 6, 4, 2, 1, 2, 3, 5, 6, 7}
	fmt.Println(testData)

	for i := 0; i < len(testData); i++ {
		//fmt.Println(i)
		for j := 0; j < i-1; j++ {
			if testData[j] > testData[j+1] {
				temp := testData[j]
				testData[j] = testData[j+1]
				testData[j+1] = temp
			}

		}
	}
	fmt.Println(testData)

}

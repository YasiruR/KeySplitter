package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(subtractN(3, 7))
	//fmt.Println(subtractN(89, 117))
	//fmt.Println(subtractN(21, 93))
	//fmt.Println(subtractN(3, 125))
	//fmt.Println(subtractN(43, 9561)) // 1582

	fmt.Println(subtract([]int{3, 89, 21, 3, 43}, []int{7, 117, 93, 125, 9561}))
}

func subtract(a, b []int) (c []int) {
	if len(a) != len(b) {
		return nil
	}

	for i := 0; i < len(a); i++ {
		c = append(c, subtractN(a[i], b[i]))
	}

	return
}

func subtractN(a, b int) int {
	var cStr string
	aStr, bStr := strconv.Itoa(a), strconv.Itoa(b)

	var excDigs int
	if len(aStr) < len(bStr) {
		excDigs = len(bStr) - len(aStr)
		for i := 0; i < excDigs; i++ {
			digB, _ := strconv.Atoi(string(bStr[i]))
			cStr += strconv.Itoa(10 - digB)
		}
	}

	for index, _ := range aStr {
		digA, _ := strconv.Atoi(string(aStr[index]))
		digB, _ := strconv.Atoi(string(bStr[excDigs+index]))

		digC := subtractD(digA, digB)
		//fmt.Printf("%d - %d = %d\n", digA, digB, digC)
		cStr += strconv.Itoa(digC)
	}

	c, _ := strconv.Atoi(cStr)
	return c
}

func subtractD(digA, digB int) (digC int) {
	if digA >= digB {
		return digA - digB
	}

	return digA + 10 - digB
}

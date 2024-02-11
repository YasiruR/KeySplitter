package main

import (
	"fmt"
	"github.com/YasiruR/keySplitter/domain"
	"github.com/YasiruR/keySplitter/services"
	"strconv"
)

func main() {
	s := services.NewSplitter()
	var secret domain.Secret = []byte(`testing 123 in ~!@#$%`)
	shares := s.Split(secret, 2)
	fmt.Println(`shares: `, shares)

	m := services.NewMerger()
	fmt.Println(`res secret: `, string(m.Merge(shares)))
}

func addN(a, b int) int {
	var cStr string
	aStr, bStr := strconv.Itoa(a), strconv.Itoa(b)

	// Adjusting numbers into same lengths by adding leading zeros
	var excDigits int
	if len(aStr) >= len(bStr) {
		excDigits = len(aStr) - len(bStr)
		for i := 0; i < excDigits; i++ {
			bStr = `0` + bStr
		}
	} else {
		excDigits = len(bStr) - len(aStr)
		for i := 0; i < excDigits; i++ {
			aStr = `0` + aStr
		}
	}

	for i, _ := range aStr {
		digA, _ := strconv.Atoi(string(aStr[i]))
		digB, _ := strconv.Atoi(string(bStr[i]))
		cStr += strconv.Itoa(addD(digA, digB))
	}

	c, _ := strconv.Atoi(cStr)
	return c
}

func addD(a, b int) int {
	sum := a + b
	if sum > 9 {
		return sum - 10
	}
	return sum
}

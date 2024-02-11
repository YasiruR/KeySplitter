package services

import (
	"github.com/YasiruR/keySplitter/domain"
	"github.com/tryfix/log"
	"math/rand"
	"strconv"
	"time"
)

type Splitter struct {
	log log.Logger
}

func NewSplitter() domain.Splitter {
	return &Splitter{
		log: log.Constructor.Log(
			log.WithColors(true),
			log.WithLevel("DEBUG"),
			log.WithFilePath(true),
		),
	}
}

// Split returns n number of shares which can be used to formulate the
// given secret when merged all together
func (s *Splitter) Split(sec domain.Secret, n int) []domain.Share {
	rand.Seed(time.Now().Unix())
	lenSec := len(sec)
	//var shares []domain.Share
	var shares [][]int

	if lenSec == 0 {
		s.log.Error(`empty secret`)
		return nil
	}

	// Convert secret into an int array
	var secIntList []int
	for _, num := range sec {
		secIntList = append(secIntList, int(num))
	}

	// Generate n-1 random int arrays with the same length
	for i := 0; i < n-1; i++ {
		// Construct a share with random numbers
		share := make([]int, lenSec)
		for j := 0; j < lenSec; j++ {
			randInt := rand.Intn(1000)
			share = append(share, randInt)
		}

		shares = append(shares, share)
	}

	// In the first iteration, first random int array is subtracted from the
	// secret and the resultant array is carried out to the next iteration. In
	// the following iterations, corresponding int arrays are subtracted from
	// the resultant arrays.
	var res []int
	var shareList []domain.Share
	for i := 0; i < n-1; i++ {
		if i == 0 {
			res = secIntList
		}

		res = s.subtract(res, shares[i])
		shareList = append(shareList, domain.Share{Id: i, PartKey: res})
	}

	// Return the array of shares
	return shareList
}

func (s *Splitter) subtract(a, b []int) (c []int) {
	if len(a) != len(b) {
		return nil
	}

	for i := 0; i < len(a); i++ {
		c = append(c, s.subtractN(a[i], b[i]))
	}

	return
}

func (s *Splitter) subtractN(a, b int) int {
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

		digC := s.subtractD(digA, digB)
		//fmt.Printf("%d - %d = %d\n", digA, digB, digC)
		cStr += strconv.Itoa(digC)
	}

	c, _ := strconv.Atoi(cStr)
	return c
}

// Each digit of an int array is subtracted from the corresponding digit in
// the second array without the carry
// (eg: 3-7 = 13-7 = 6, 89-117 = 972, 21-93 = 38, 3-125 = 988)
func (s *Splitter) subtractD(digA, digB int) (digC int) {
	if digA >= digB {
		return digA - digB
	}

	return digA + 10 - digB
}

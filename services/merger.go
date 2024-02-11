package services

import (
	"github.com/YasiruR/keySplitter/domain"
	"github.com/tryfix/log"
	"sort"
	"strconv"
)

type Merger struct {
	log log.Logger
}

func NewMerger() domain.Merger {
	return &Merger{
		log: log.Constructor.Log(
			log.WithColors(true),
			log.WithLevel("DEBUG"),
			log.WithFilePath(true),
		),
	}
}

func (m *Merger) Merge(shares []domain.Share) domain.Secret {
	// Sort share list based on IDs in the descending order
	// todo CHECK SORT
	sort.Slice(shares, func(i, j int) bool {
		return shares[i].Id > shares[j].Id
	})

	var res []int
	for i, share := range shares {
		if i == 0 {
			res = m.add(shares[i].PartKey, shares[i+1].PartKey)
			continue
		}

		if share.Id == 0 || i == len(shares)-1 {
			break
		}

		res = m.add(res, shares[i+1].PartKey)
	}

	// Convert final sum into a slice of bytes
	var cByts []byte
	for _, n := range res {
		cByts = append(cByts, byte(n))
	}

	return cByts
}

func (m *Merger) add(a, b []int) (c []int) {
	if len(a) != len(b) {
		return nil
	}

	for i, _ := range a {
		c = append(c, m.addN(a[i], b[i]))
	}

	return
}

func (m *Merger) addN(a, b int) int {
	var cStr string
	aStr, bStr := numToSameLen(strconv.Itoa(a), strconv.Itoa(b))

	for i, _ := range aStr {
		digA, _ := strconv.Atoi(string(aStr[i]))
		digB, _ := strconv.Atoi(string(bStr[i]))
		cStr += strconv.Itoa(m.addD(digA, digB))
	}

	c, _ := strconv.Atoi(cStr)
	return c
}

func (m *Merger) addD(a, b int) int {
	sum := a + b
	if sum > 9 {
		return sum - 10
	}
	return sum
}

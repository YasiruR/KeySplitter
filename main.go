package main

import (
	"fmt"
	"github.com/YasiruR/keySplitter/domain"
	"github.com/YasiruR/keySplitter/services"
)

func main() {
	s := services.NewSplitter()
	var secret domain.Secret = []byte(`testing 123 in ~!@#$%`)
	shares := s.Split(secret, 2)
	fmt.Println(`shares: `, shares)

	m := services.NewMerger()
	fmt.Println(`res secret: `, string(m.Merge(shares)))
}

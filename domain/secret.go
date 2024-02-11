package domain

type Secret []byte

type PartKey []int

// Share is PartKey with the index
type Share struct {
	Id int
	PartKey
}

type PartKeyService interface {
	Bytes()
	Parse()
}

type Splitter interface {
	Split(s Secret, n int) []Share
}

type Merger interface {
	Merge([]Share) Secret
}

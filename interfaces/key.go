package interfaces

type Key []byte

type PartKey []byte

// Share is PartKey with the index
type Share map[int]PartKey

type KeyService interface {
	Init()
	Split()
}

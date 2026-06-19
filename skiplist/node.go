package skiplist

type Node struct {
	Key string

	Entry Entry

	Forward []*Node
}
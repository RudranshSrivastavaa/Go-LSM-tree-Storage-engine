package skiplist

type Entry struct {
	Value     []byte
	Tombstone bool
}
package skiplist

import (
	"math/rand"
	"time"
)

const (
	MaxLevel = 16
	P        = 0.5
)

type SkipList struct {
	level int
	head  *Node
	rand  *rand.Rand
}

func New() *SkipList {

	head := &Node{
		Forward: make([]*Node, MaxLevel),
	}

	return &SkipList{
		level: 1,
		head:  head,
		rand: rand.New(
			rand.NewSource(
				time.Now().UnixNano(),
			),
		),
	}
}

func (s *SkipList) randomLevel() int {

	level := 1

	for s.rand.Float64() < P &&
		level < MaxLevel {

		level++
	}

	return level
}

func (s *SkipList) Insert(
	key string,
	entry Entry,
) {

	update := make([]*Node, MaxLevel)

	current := s.head

	for i := s.level - 1; i >= 0; i-- {

		for current.Forward[i] != nil &&
			current.Forward[i].Key < key {

			current = current.Forward[i]
		}

		update[i] = current
	}

	current = current.Forward[0]

	if current != nil &&
		current.Key == key {

		current.Entry = entry
		return
	}

	level := s.randomLevel()

	if level > s.level {

		for i := s.level; i < level; i++ {
			update[i] = s.head
		}

		s.level = level
	}

	node := &Node{
		Key:     key,
		Entry:   entry,
		Forward: make([]*Node, level),
	}

	for i := 0; i < level; i++ {

		node.Forward[i] = update[i].Forward[i]

		update[i].Forward[i] = node
	}
}

func (s *SkipList) Search(
	key string,
) (*Entry, bool) {

	current := s.head

	for i := s.level - 1; i >= 0; i-- {

		for current.Forward[i] != nil &&
			current.Forward[i].Key < key {

			current = current.Forward[i]
		}
	}

	current = current.Forward[0]

	if current != nil &&
		current.Key == key {

		return &current.Entry, true
	}

	return nil, false
}

func (s *SkipList) Delete(
	key string,
) {

	entry := Entry{
		Tombstone: true,
	}

	s.Insert(
		key,
		entry,
	)
}

func (s *SkipList) RangeScan(
	start,
	end string,
) []Node {

	var result []Node

	current := s.head

	for i := s.level - 1; i >= 0; i-- {

		for current.Forward[i] != nil &&
			current.Forward[i].Key < start {

			current = current.Forward[i]
		}
	}

	current = current.Forward[0]

	for current != nil &&
		current.Key <= end {

		if !current.Entry.Tombstone {
			result = append(
				result,
				*current,
			)
		}

		current = current.Forward[0]
	}

	return result
}

func (s *SkipList) Iterate(fn func(string, Entry) bool) {

    current := s.head.Forward[0]

    for current != nil {

        if !fn(
            current.Key,
            current.Entry,
        ) {
            return
        }

        current =
            current.Forward[0]
    }
}

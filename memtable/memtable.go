package memtable

import "storage-engine/skiplist"

type MemTable struct {
    SkipList *skiplist.SkipList
    Size     int64
}

func New() *MemTable {
    return &MemTable{
        SkipList: skiplist.New(),
    }
}

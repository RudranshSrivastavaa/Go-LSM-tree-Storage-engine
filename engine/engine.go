package engine

import (
	"fmt"
	"storage-engine/memtable"
	"storage-engine/skiplist"
	"storage-engine/sstable"
	"storage-engine/wal"
)

const MemTableThreshold = 1024

type Engine struct {
	active    *memtable.MemTable
	immutable *memtable.MemTable

	wal     *wal.WAL
	nextSST uint64
	manager *sstable.Manager
}

func New(path string) (*Engine, error) {

	recoveredSkipList, err := wal.Recover(path)

	if err != nil {
		return nil, err
	}

	log, err := wal.New(path)

	if err != nil {
		return nil, err
	}

	manager, err := sstable.LoadManager()

	if err != nil {
		return nil, err
	}

	active := &memtable.MemTable{
		SkipList: recoveredSkipList,
	}

	return &Engine{
		active:  active,
		manager: manager,
		wal:     log,
	}, nil
}

func (e *Engine) Put(key string, value []byte) error {

	record := fmt.Sprintf(
		"PUT|%s|%s",
		key,
		string(value),
	)

	if err := e.wal.Write(record); err != nil {
		return err
	}

	e.active.SkipList.Insert(
		key,
		skiplist.Entry{
			Value: value,
		},
	)

	e.active.Size += EstimateEntrySize(
		key,
		value,
	)

	if e.active.Size >= MemTableThreshold {

		e.rotateMemTable()
	}

	return nil
}

func (e *Engine) Get(key string) ([]byte, bool) {

	// Active

	if entry, ok :=
		e.active.SkipList.Search(key); ok {

		if entry.Tombstone {
			return nil, false
		}

		return entry.Value, true
	}

	if e.immutable != nil {

		if entry, ok :=
			e.immutable.SkipList.Search(key); ok {

			if entry.Tombstone {
				return nil, false
			}

			return entry.Value, true
		}
	}
	files := e.manager.Files()

	for i := len(files) - 1; i >= 0; i-- {

		reader :=
			sstable.NewReader(
				files[i],
			)
			if reader.Bloom() != nil {

	if !reader.Bloom().MayContain(
		key,
	) {
		continue
	}
}
		rec,
			found,
			err :=
			reader.Get(key)

		if err != nil {
			continue
		}

		if found {

			if rec.Tombstone {
				return nil, false
			}

			return rec.Value, true
		}
	}

	return nil, false
}

func (e *Engine) Delete(key string) error {

	record := fmt.Sprintf(
		"DEL|%s",
		key,
	)

	if err := e.wal.Write(record); err != nil {
		return err
	}

	e.active.SkipList.Delete(key)

	e.active.Size += EstimateEntrySize(
		key,
		nil,
	)

	if e.active.Size >= MemTableThreshold {
		e.rotateMemTable()
	}

	return nil
}

func (e *Engine) Close() error {
	return e.wal.Close()
}

func EstimateEntrySize(
	key string,
	value []byte,
) int64 {

	return int64(
		len(key) +
			len(value) +
			16,
	)
}

func (e *Engine) rotateMemTable() {
	e.immutable = e.active

	sstPath := fmt.Sprintf(
		"data/sst_%06d.db",
		e.nextSST,
	)

	e.nextSST++

	err := sstable.Flush(
		e.immutable,
		sstPath,
	)
	if err == nil {
		e.manager.Add(sstPath)
	}

	e.active = memtable.New()

	e.immutable = nil
}

func (e *Engine) Active() *memtable.MemTable {
	return e.active
}

func (e *Engine) Immutable() *memtable.MemTable {
	return e.immutable
}

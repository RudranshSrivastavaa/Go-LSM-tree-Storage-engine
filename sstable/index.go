package sstable

import (
	"encoding/binary"
	"os"
	"sort"
)

func writeIndex(
	file *os.File,
	index []IndexEntry,
) error {

	count := uint64(len(index))

	if err := binary.Write(
		file,
		binary.LittleEndian,
		count,
	); err != nil {
		return err
	}

	for _, entry := range index {

		keyLen := uint32(
			len(entry.Key),
		)

		if err := binary.Write(
			file,
			binary.LittleEndian,
			keyLen,
		); err != nil {
			return err
		}

		if _, err := file.Write(
			[]byte(entry.Key),
		); err != nil {
			return err
		}

		if err := binary.Write(
			file,
			binary.LittleEndian,
			entry.Offset,
		); err != nil {
			return err
		}
	}

	return nil
}

func (r *Reader) findOffset(
	key string,
) int64 {

	if len(r.index) == 0 {
		return 0
	}

	pos := sort.Search(
		len(r.index),
		func(i int) bool {
			return r.index[i].Key > key
		},
	)

	if pos == 0 {
		return 0
	}

	return r.index[pos-1].Offset
}
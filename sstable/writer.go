package sstable

import (
	"encoding/binary"
	"io"
	"os"
	"fmt"
	"storage-engine/bloom"
	"storage-engine/memtable"
	"storage-engine/skiplist"
)

func writeRecord(
	file *os.File,
	r Record,
) error {

	keyLen := uint32(len(r.Key))
	valueLen := uint32(len(r.Value))

	if err := binary.Write(
		file,
		binary.LittleEndian,
		keyLen,
	); err != nil {
		return err
	}

	if err := binary.Write(
		file,
		binary.LittleEndian,
		valueLen,
	); err != nil {
		return err
	}

	var tombstone uint8

	if r.Tombstone {
		tombstone = 1
	}

	if err := binary.Write(
		file,
		binary.LittleEndian,
		tombstone,
	); err != nil {
		return err
	}

	if _, err := file.Write(
		[]byte(r.Key),
	); err != nil {
		return err
	}

	if _, err := file.Write(
		r.Value,
	); err != nil {
		return err
	}

	return nil
}

func Flush(
	mem *memtable.MemTable,
	path string,
) error {

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	var count uint64

	mem.SkipList.Iterate(
		func(
			key string,
			entry skiplist.Entry,
		) bool {

			count++
			return true
		},
	)

	header := Header{
		Version:     1,
		RecordCount: count,
		IndexOffset: 0,
		BloomOffset: 0,
	}

	bf := bloom.New(count)

	// Placeholder header
	if err := writeHeader(
		file,
		header,
	); err != nil {
		return err
	}

	var index []IndexEntry

	recordNum := 0

	mem.SkipList.Iterate(
		func(
			key string,
			entry skiplist.Entry,
		) bool {

			offset, _ := file.Seek(
				0,
				io.SeekCurrent,
			)

			if recordNum%IndexInterval == 0 {

				index = append(
					index,
					IndexEntry{
						Key:    key,
						Offset: offset,
					},
				)
			}

			recordNum++

			bf.Add(key)

			err := writeRecord(
				file,
				Record{
					Key:       key,
					Value:     entry.Value,
					Tombstone: entry.Tombstone,
				},
			)

			return err == nil
		},
	)

	// Index starts here
	indexOffset, _ := file.Seek(
		0,
		io.SeekCurrent,
	)

	header.IndexOffset = indexOffset

	err = writeIndex(
		file,
		index,
	)

	if err != nil {
		return err
	}

	// Bloom starts here
	bloomOffset, _ := file.Seek(
		0,
		io.SeekCurrent,
	)

	header.BloomOffset = bloomOffset

	err = writeBloom(
		file,
		bf,
	)

	if err != nil {
		return err
	}

	// DEBUG
	fileSize, _ := file.Seek(
		0,
		io.SeekCurrent,
	)

	fmt.Println("================================")
	fmt.Println("FINAL HEADER VALUES")
	fmt.Println("RecordCount :", header.RecordCount)
	fmt.Println("IndexOffset :", header.IndexOffset)
	fmt.Println("BloomOffset :", header.BloomOffset)
	fmt.Println("File Size   :", fileSize)
	fmt.Println("================================")

	// Go back and rewrite header
	_, err = file.Seek(
		0,
		io.SeekStart,
	)

	if err != nil {
		return err
	}

	err = writeHeader(
		file,
		header,
	)

	if err != nil {
		return err
	}

	fmt.Println("HEADER REWRITTEN")

	return file.Sync()
}

func writeHeader(
	file *os.File,
	header Header,
) error {

	if err := binary.Write(
		file,
		binary.LittleEndian,
		header.Version,
	); err != nil {
		return err
	}

	if err := binary.Write(
		file,
		binary.LittleEndian,
		header.RecordCount,
	); err != nil {
		return err
	}

	if err := binary.Write(
		file,
		binary.LittleEndian,
		header.IndexOffset,
	); err != nil {
		return err
	}
	if err := binary.Write(
		file,
		binary.LittleEndian,
		header.BloomOffset,
	); err != nil {
		return err
	}

	return nil
}

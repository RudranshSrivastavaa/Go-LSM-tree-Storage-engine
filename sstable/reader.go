package sstable

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
    "storage-engine/bloom"
)

type Reader struct {
	path string

	index []IndexEntry
    bloom *bloom.Filter
}

func NewReader(
	path string,
) *Reader {

	r := &Reader{
		path: path,
	}

	file, err := os.Open(path)

	if err != nil {
		return r
	}

	defer file.Close()

	header, err := ReadHeader(file)

	if err != nil {
		return r
	}
    fmt.Println(
	"Version:",
	header.Version,
)

fmt.Println(
	"RecordCount:",
	header.RecordCount,
)

fmt.Println(
	"IndexOffset:",
	header.IndexOffset,
)

fmt.Println(
	"BloomOffset:",
	header.BloomOffset,
)

	index, err := loadIndex(
		file,
		header.IndexOffset,
	)

	if err == nil {
		r.index = index
	}

	bf, err := loadBloom(
		file,
		header.BloomOffset,
	)

fmt.Println(
	"BloomOffset:",
	header.BloomOffset,
)

fmt.Println(
	"Bloom Load Error:",
	err,
)

	if err == nil {
		r.bloom = bf
	}

	return r
}

func readRecord(file *os.File) (*Record, error) {

	var keyLen uint32
	var valueLen uint32
	var tombstone uint8

	err := binary.Read(
		file,
		binary.LittleEndian,
		&keyLen,
	)

	if err != nil {
		return nil, err
	}

	err = binary.Read(
		file,
		binary.LittleEndian,
		&valueLen,
	)

	if err != nil {
		return nil, err
	}

	err = binary.Read(
		file,
		binary.LittleEndian,
		&tombstone,
	)

	if err != nil {
		return nil, err
	}

	key := make([]byte, keyLen)

	_, err = io.ReadFull(
		file,
		key,
	)

	if err != nil {
		return nil, err
	}

	value := make([]byte, valueLen)

	_, err = io.ReadFull(
		file,
		value,
	)

	if err != nil {
		return nil, err
	}

	return &Record{
		Key:       string(key),
		Value:     value,
		Tombstone: tombstone == 1,
	}, nil
}

func (r *Reader) Get(
	key string,
) (*Record, bool, error) {

	file, err := os.Open(r.path)

	if err != nil {
		return nil, false, err
	}

	defer file.Close()

	var header Header

	err = binary.Read(
		file,
		binary.LittleEndian,
		&header.Version,
	)

	if err != nil {
		return nil, false, err
	}

	err = binary.Read(
		file,
		binary.LittleEndian,
		&header.RecordCount,
	)

	if err != nil {
		return nil, false, err
	}

	err = binary.Read(
		file,
		binary.LittleEndian,
		&header.IndexOffset,
	)

	if err != nil {
		return nil, false, err
	}
    err = binary.Read(
		file,
		binary.LittleEndian,
		&header.BloomOffset,
	)

	if err != nil {
		return nil, false, err
	}

	offset := r.findOffset(key)

	_, err = file.Seek(
		offset,
		io.SeekStart,
	)

	if err != nil {
		return nil, false, err
	}

	for {

		pos, err := file.Seek(
			0,
			io.SeekCurrent,
		)

		if err != nil {
			return nil, false, err
		}

		if pos >= header.IndexOffset {
			break
		}

		rec, err := readRecord(file)

		if err != nil {
			return nil, false, err
		}

		if rec.Key == key {
			return rec, true, nil
		}

		if rec.Key > key {
			return nil, false, nil
		}
	}

	return nil, false, nil
}
func (r *Reader) LinearGet(
	key string,
) (*Record, bool, error) {

	file, err := os.Open(r.path)

	if err != nil {
		return nil, false, err
	}

	defer file.Close()

	var header Header

	err = binary.Read(
		file,
		binary.LittleEndian,
		&header.Version,
	)

	if err != nil {
		return nil, false, err
	}

	err = binary.Read(
		file,
		binary.LittleEndian,
		&header.RecordCount,
	)

	if err != nil {
		return nil, false, err
	}

	err = binary.Read(
		file,
		binary.LittleEndian,
		&header.IndexOffset,
	)

	if err != nil {
		return nil, false, err
	}
	err = binary.Read(
	file,
	binary.LittleEndian,
	&header.BloomOffset,
)

if err != nil {
	return nil, false, err
}
	for i := uint64(0); i < header.RecordCount; i++ {

		rec, err := readRecord(file)

		if err != nil {
			return nil, false, err
		}

		if rec.Key == key {
			return rec, true, nil
		}
	}

	return nil, false, nil
}
func ReadHeader(
	file *os.File,
) (Header, error) {

	var header Header

	err := binary.Read(
		file,
		binary.LittleEndian,
		&header.Version,
	)

	if err != nil {
		return header, err
	}

	err = binary.Read(
		file,
		binary.LittleEndian,
		&header.RecordCount,
	)

	if err != nil {
		return header, err
	}

	err = binary.Read(
		file,
		binary.LittleEndian,
		&header.IndexOffset,
	)

	if err != nil {
		return header, err
	}
	err = binary.Read(
		file,
		binary.LittleEndian,
		&header.BloomOffset,
	)
	if err != nil {
		return header, err
	}

	return header, nil
}

func loadIndex(
	file *os.File,
	offset int64,
) ([]IndexEntry, error) {

	_, err := file.Seek(
		offset,
		io.SeekStart,
	)

	if err != nil {
		return nil, err
	}

	var count uint64

	err = binary.Read(
		file,
		binary.LittleEndian,
		&count,
	)
	fmt.Println(
		"Index Count:",
		count,
	)

	if err != nil {
		return nil, err
	}

	index := make(
		[]IndexEntry,
		0,
		count,
	)

	for i := uint64(0); i < count; i++ {

		var keyLen uint32

		err := binary.Read(
			file,
			binary.LittleEndian,
			&keyLen,
		)

		if err != nil {
			return nil, err
		}

		key := make(
			[]byte,
			keyLen,
		)

		_, err = io.ReadFull(
			file,
			key,
		)

		if err != nil {
			return nil, err
		}

		var offset int64

		err = binary.Read(
			file,
			binary.LittleEndian,
			&offset,
		)

		if err != nil {
			return nil, err
		}

		index = append(
			index,
			IndexEntry{
				Key:    string(key),
				Offset: offset,
			},
		)
	}

	return index, nil
}

func (r *Reader) IndexCount() int {
	return len(r.index)
}

func (r *Reader) Bloom() *bloom.Filter {
	return r.bloom
}
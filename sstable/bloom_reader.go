package sstable

import (
	"encoding/binary"
	"io"
	"os"

	"storage-engine/bloom"
)

func loadBloom(
	file *os.File,
	offset int64,
) (*bloom.Filter, error) {

	_, err := file.Seek(
		offset,
		io.SeekStart,
	)

	if err != nil {
		return nil, err
	}

	var size uint64

	err = binary.Read(
		file,
		binary.LittleEndian,
		&size,
	)

	if err != nil {
		return nil, err
	}

	bits := make(
		[]byte,
		size,
	)

	_, err = io.ReadFull(
		file,
		bits,
	)

	if err != nil {
		return nil, err
	}

	return bloom.FromBytes(
		bits,
	), nil
}
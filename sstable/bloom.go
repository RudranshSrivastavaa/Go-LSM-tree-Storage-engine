package sstable

import (
	"encoding/binary"
	"os"

	"storage-engine/bloom"
)

func writeBloom(
	file *os.File,
	filter *bloom.Filter,
) error {

	size := uint64(
		len(filter.Bits()),
	)

	err := binary.Write(
		file,
		binary.LittleEndian,
		size,
	)

	if err != nil {
		return err
	}

	_, err = file.Write(
		filter.Bits(),
	)

	return err
}
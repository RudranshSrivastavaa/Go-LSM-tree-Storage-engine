package main

import (
	"fmt"
	"os"

	"storage-engine/sstable"
)

func main() {

	file, err := os.Open(
		"data/sst_000000.db",
	)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	header, err := sstable.ReadHeader(
		file,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"Index:",
		header.IndexOffset,
	)

	fmt.Println(
		"Bloom:",
		header.BloomOffset,
	)
}
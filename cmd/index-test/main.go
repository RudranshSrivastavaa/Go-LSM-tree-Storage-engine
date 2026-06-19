package main

import (
	"fmt"
	"storage-engine/sstable"
)

func main() {

	reader := sstable.NewReader(
		"data/sst_000001.db",
	)

	fmt.Println(
		"Index Entries:",
		reader.IndexCount(),
	)
}
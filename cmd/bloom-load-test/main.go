package main

import (
	"fmt"

	"storage-engine/sstable"
)

func main() {

	reader := sstable.NewReader(
		"data/sst_000000.db",
	)

	if reader.Bloom() == nil {
		fmt.Println("Bloom filter not loaded")
		return
	}

	fmt.Println(
		reader.Bloom().MayContain(
			"k0001",
		),
	)

	fmt.Println(
		reader.Bloom().MayContain(
			"does-not-exist",
		),
	)
}
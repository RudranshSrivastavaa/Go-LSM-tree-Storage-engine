package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open(
		"data/sst_000001.db",
	)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var version uint32
	var count uint64

	err = binary.Read(
		file,
		binary.LittleEndian,
		&version,
	)

	if err != nil {
		panic(err)
	}

	err = binary.Read(
		file,
		binary.LittleEndian,
		&count,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"Version:",
		version,
	)

	fmt.Println(
		"RecordCount:",
		count,
	)
}
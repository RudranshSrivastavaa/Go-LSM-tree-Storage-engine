package main

import (
	"fmt"
	"storage-engine/wal"
)

func main() {

	log, _ := wal.New("large.log")
	defer log.Close()

	for i := 0; i < 10000; i++ {

		log.Write(
			fmt.Sprintf(
				"PUT|k%d|value",
				i,
			),
		)
	}
}
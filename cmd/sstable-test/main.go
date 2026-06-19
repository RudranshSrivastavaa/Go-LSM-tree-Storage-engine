package main

import (
	"fmt"

	"storage-engine/engine"
)

func main() {

	db, err := engine.New("wal.log")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 500; i++ {

		err := db.Put(
			fmt.Sprintf("k%03d", i),
			[]byte("value"),
		)

		if err != nil {
			panic(err)
		}
	}

	fmt.Println("done")
}
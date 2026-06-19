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

	for i := 0; i < 5000; i++ {

		err := db.Put(
			fmt.Sprintf(
				"k%04d",
				i,
			),
			[]byte("value"),
		)

		if err != nil {
			panic(err)
		}
	}

	db.Close()

	fmt.Println("Data written")

	db, err = engine.New("wal.log")
	if err != nil {
		panic(err)
	}

	keys := []string{
		"k0001",
		"k1000",
		"k3000",
		"k4999",
	}

	for _, k := range keys {

		v, ok := db.Get(k)

		fmt.Println(
			k,
			ok,
			string(v),
		)
	}
}
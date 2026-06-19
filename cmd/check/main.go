package main

import (
	"fmt"
	"storage-engine/engine"
)

func main() {

	db, err := engine.New("large.log")
	if err != nil {
		panic(err)
	}

	keys := []string{
		"k0",
		"k1",
		"k100",
		"k5000",
		"k9999",
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
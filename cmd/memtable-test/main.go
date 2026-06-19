package main

import (
	"fmt"
	"storage-engine/engine"
)

func main() {

	db, err := engine.New("test.log")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 250; i++ {

		err := db.Put(
			fmt.Sprintf("k%d", i),
			[]byte("value"),
		)

		if err != nil {
			panic(err)
		}
	}

	value, ok := db.Get("k10")

	fmt.Println(
		string(value),
		ok,
	)

	fmt.Println(
		"active exists:",
		db.Active() != nil,
	)

	fmt.Println(
		"immutable exists:",
		db.Immutable() != nil,
	)
}
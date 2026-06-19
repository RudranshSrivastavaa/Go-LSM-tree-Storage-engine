package main

import (
	"fmt"
	"storage-engine/engine"
)

func main() {

	db, _ := engine.New("data/wal.log")

	defer db.Close()

	db.Put("user1", []byte("alice"))
	db.Put("user2", []byte("bob"))

	db.Delete("user1")

	value, ok := db.Get("user2")

	fmt.Println(
		string(value),
		ok,
	)
}
package main

import (
	"fmt"

	"storage-engine/bloom"
)

func main() {

	bf := bloom.New(1000)

	bf.Add("alice")

	bf.Add("bob")

	fmt.Println(
		bf.MayContain("alice"),
	)

	fmt.Println(
		bf.MayContain("bob"),
	)

	fmt.Println(
		bf.MayContain("charlie"),
	)
}
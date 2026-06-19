package bloom

import (
	"hash/fnv"
)

func hash1(
	key string,
) uint64 {

	h := fnv.New64a()

	h.Write(
		[]byte(key),
	)

	return h.Sum64()
}

func hash2(
	key string,
) uint64 {

	h := fnv.New64()

	h.Write(
		[]byte(key),
	)

	return h.Sum64()
}
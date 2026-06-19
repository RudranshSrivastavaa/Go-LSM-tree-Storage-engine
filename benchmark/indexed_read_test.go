package benchmark

import (
	"testing"

	"storage-engine/sstable"
)

func BenchmarkLinearRead(
	b *testing.B,
) {

	reader := sstable.NewReader(
	"../data/sst_000001.db",
)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_, _, err :=
			reader.LinearGet(
				"k0500",
			)

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIndexedRead(
	b *testing.B,
) {

	reader := sstable.NewReader(
	"../data/sst_000001.db",
)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_, _, err :=
			reader.Get(
				"k0500",
			)

		if err != nil {
			b.Fatal(err)
		}
	}
}
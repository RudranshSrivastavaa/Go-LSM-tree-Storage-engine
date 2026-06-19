package benchmark

import (
	"testing"

	"storage-engine/sstable"
)

func BenchmarkNegativeLookup(
	b *testing.B,
) {

	reader := sstable.NewReader(
	"../data/sst_000000.db",
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		reader.Get(
			"does-not-exist",
		)
	}
}

func BenchmarkBloomNegativeLookup(
	b *testing.B,
) {

	reader := sstable.NewReader(
	"../data/sst_000001.db",
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		if !reader.Bloom().MayContain(
			"does-not-exist",
		) {
			continue
		}

		reader.Get(
			"does-not-exist",
		)
	}
}
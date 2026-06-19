package benchmark

import (
	"testing"

	"storage-engine/memtable"
)

func BenchmarkMemTableSwitch(
	b *testing.B,
) {

	for i := 0; i < b.N; i++ {

		active := memtable.New()

		immutable := active

		_ = immutable

		active = memtable.New()

		_ = active
	}
}
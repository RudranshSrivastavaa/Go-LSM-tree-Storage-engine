package benchmark

import (
	"fmt"
	"testing"

	"storage-engine/memtable"
	"storage-engine/skiplist"
	"storage-engine/sstable"
)

func BenchmarkFlush(
	b *testing.B,
) {

	mem := memtable.New()

	for i := 0; i < 100000; i++ {

		mem.SkipList.Insert(
			fmt.Sprintf("k%08d", i),
			skiplist.Entry{
				Value: []byte("value"),
			},
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		err := sstable.Flush(
			mem,
			"test.sst",
		)

		if err != nil {
			b.Fatal(err)
		}
	}
}
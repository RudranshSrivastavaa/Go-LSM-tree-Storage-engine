package benchmark

import (
	"fmt"
	"testing"

	"storage-engine/engine"
)

func buildLargeDB(
	b *testing.B,
) *engine.Engine {

	db, err := engine.New(
		"bench_read.log",
	)

	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < 10000; i++ {

		err := db.Put(
			fmt.Sprintf(
				"k%06d",
				i,
			),
			[]byte("value"),
		)

		if err != nil {
			b.Fatal(err)
		}
	}

	return db
}

func BenchmarkRead(
	b *testing.B,
) {

	db := buildLargeDB(b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_, _ = db.Get(
			"k5000",
		)
	}
}
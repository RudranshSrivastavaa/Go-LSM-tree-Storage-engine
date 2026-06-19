package benchmark

import (
	"testing"
	"fmt"
	"storage-engine/engine"

)

func BenchmarkRecovery(
	b *testing.B,
) {

	db, err := engine.New(
		"recovery.log",
	)

	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < 10000; i++ {

		db.Put(
			fmt.Sprintf(
				"k%d",
				i,
			),
			[]byte("value"),
		)
	}

	db.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_, err := engine.New(
			"recovery.log",
		)

		if err != nil {
			b.Fatal(err)
		}
	}
}
package benchmark

import (
	"fmt"
	"os"
	"testing"

	"storage-engine/engine"
)

func BenchmarkPut(b *testing.B) {

	os.Remove("bench.log")

	db, err := engine.New("bench.log")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		err := db.Put(
			fmt.Sprintf("k%d", i),
			[]byte("value"),
		)

		if err != nil {
			b.Fatal(err)
		}
	}
}
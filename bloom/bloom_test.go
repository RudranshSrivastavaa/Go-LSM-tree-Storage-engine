package bloom

import (
	"fmt"
	"testing"
)

func BenchmarkMayContain(
	b *testing.B,
) {

	filter := New(
		100000,
	)

	for i := 0; i < 100000; i++ {

		filter.Add(
			fmt.Sprintf(
				"k%d",
				i,
			),
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		filter.MayContain(
			"k50000",
		)
	}
}
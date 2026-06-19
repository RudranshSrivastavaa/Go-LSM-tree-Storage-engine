package skiplist

import(
	"fmt"
	"testing"
)

func BenchmarkInsert(
	b *testing.B,
) {

	sl := New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		sl.Insert(
			fmt.Sprintf("k%08d", i),
			Entry{
				Value: []byte("value"),
			},
		)
	}
}

func BenchmarkSearch(
	b *testing.B,
) {

	sl := New()

	for i := 0; i < 1000000; i++ {

		sl.Insert(
			fmt.Sprintf("k%08d", i),
			Entry{
				Value: []byte("value"),
			},
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		sl.Search("k00050000")
	}
}
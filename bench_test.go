package fenwick

import (
	"testing"
)

func BenchmarkNew(b *testing.B) {
	n := 1000
	b.StopTimer()
	a := make([]int64, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = New(a...)
	}
}

func BenchmarkSum(b *testing.B) {
	n := 1000
	b.StopTimer()
	a := make([]int64, n)
	l := New(a...)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			_ = l.Sum(j)
		}
	}
}

func BenchmarkSumSlice(b *testing.B) {
	n := 1000
	b.StopTimer()
	a := make([]int64, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			var sum int64
			for k := 0; k < j; k++ {
				sum += a[k]
			}
			_ = sum
		}
	}
}

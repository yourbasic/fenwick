package fenwick

import (
	"testing"
)

const n = 1000

func BenchmarkNew(b *testing.B) {
	b.StopTimer()
	a := make([]int64, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = New(a...)
	}
}

func BenchmarkAppend(b *testing.B) {
	l := new(List)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			l.Append(int64(j))
		}
	}
}

func BenchmarkSum(b *testing.B) {
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

func BenchmarkGet(b *testing.B) {
	b.StopTimer()
	a := make([]int64, n)
	l := New(a...)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			_ = l.Get(j)
		}
	}
}

func BenchmarkSet(b *testing.B) {
	b.StopTimer()
	a := make([]int64, n)
	l := New(a...)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			l.Set(j, int64(j))
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	b.StopTimer()
	a := make([]int64, n)
	l := New(a...)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			l.Add(j, int64(j))
		}
	}
}

package fenwick

import (
	"testing"
)

func TestGet(t *testing.T) {
	l := New(1, 2, 3, -4)
	if n := l.Get(0); n != 1 {
		t.Errorf("Get(0) = %d; want %d", n, 1)
	}
	if n := l.Get(1); n != 2 {
		t.Errorf("Get(1) = %d; want %d", n, 2)
	}
	if n := l.Get(2); n != 3 {
		t.Errorf("Get(2) = %d; want %d", n, 3)
	}
	if n := l.Get(3); n != -4 {
		t.Errorf("Get(3) = %d; want %d", n, -4)
	}
	n := 50
	a := make([]int64, n)
	l = New()
	for i := range a {
		a[i] = int64(i)
	}
	l = New(a...)
	for i := range a {
		if a[i] != l.Get(i) {
			t.Errorf("Get(%d) = %d; want %d", i, l.Get(i), a[i])
		}
	}
}

func TestSet(t *testing.T) {
	l := New(1, 2, 3, -4)
	l.Set(0, 4)
	if n := l.Get(0); n != 4 {
		t.Errorf("Get(0) = %d; want %d", n, 4)
	}
	l.Set(1, 5)
	if n := l.Get(1); n != 5 {
		t.Errorf("Get(1) = %d; want %d", n, 5)
	}
	l.Set(2, 6)
	if n := l.Get(0); n != 4 {
		t.Errorf("Get(0) = %d; want %d", n, 4)
	}
	if n := l.Get(1); n != 5 {
		t.Errorf("Get(1) = %d; want %d", n, 5)
	}
	if n := l.Get(2); n != 6 {
		t.Errorf("Get(2) = %d; want %d", n, 6)
	}
	if n := l.Get(3); n != -4 {
		t.Errorf("Get(3) = %d; want %d", n, -4)
	}
	n := 50
	a := make([]int64, n)
	l = New()
	for i := range a {
		a[i] = int64(i)
	}
	l = New(a...)
	for i := range a {
		l.Set(i, 100)
		if l.Get(i) != 100 {
			t.Errorf("Get(%d) = %d; want %d", i, l.Get(i), 100)
		}
	}
}

func TestAdd(t *testing.T) {
	l := New(1, 2, 3, -4)
	l.Add(0, 3)
	if n := l.Get(0); n != 4 {
		t.Errorf("Get(0) = %d; want %d", n, 4)
	}
	l.Add(1, 3)
	if n := l.Get(1); n != 5 {
		t.Errorf("Get(1) = %d; want %d", n, 5)
	}
	l.Add(2, 3)
	if n := l.Get(0); n != 4 {
		t.Errorf("Get(0) = %d; want %d", n, 4)
	}
	if n := l.Get(1); n != 5 {
		t.Errorf("Get(1) = %d; want %d", n, 5)
	}
	if n := l.Get(2); n != 6 {
		t.Errorf("Get(2) = %d; want %d", n, 6)
	}
	if n := l.Get(3); n != -4 {
		t.Errorf("Get(3) = %d; want %d", n, -4)
	}
	n := 50
	a := make([]int64, n)
	l = New()
	for i := range a {
		a[i] = int64(i)
	}
	l = New(a...)
	for i := range a {
		l.Add(i, 100)
		if l.Get(i) != a[i]+100 {
			t.Errorf("Get(%d) = %d; want %d", i, l.Get(i), a[i]+100)
		}
	}
}

func TestSum(t *testing.T) {
	l := New(1, 2, 3, -4)
	if n := l.Sum(1); n != 1 {
		t.Errorf("Sum(1) = %d; want %d", n, 1)
	}
	if n := l.Sum(2); n != 3 {
		t.Errorf("Sum(2) = %d; want %d", n, 3)
	}
	if n := l.Sum(3); n != 6 {
		t.Errorf("Sum(3) = %d; want %d", n, 6)
	}
	if n := l.Sum(4); n != 2 {
		t.Errorf("Sum(4) = %d; want %d", n, 2)
	}
	n := 50
	a := make([]int64, n)
	l = New()
	for i := range a {
		a[i] = int64(i)
	}
	l = New(a...)
	for i := range a {
		var res int64
		for j := 0; j < i; j++ {
			res += a[j]
		}
		if l.Sum(i) != res {
			t.Errorf("Sum(%d) = %d; want %d", i, l.Get(i), res)
		}
	}
}

/*
func TestAppend(t *testing.T) {
	l := New(1, 2, 3, -4)
	l.Append(5)
	if n := l.Get(4); n != 5 {
		t.Errorf("Get(4) = %d; want %d", n, 5)
	}
	if n := l.Sum(5); n != 7 {
		t.Errorf("Sum(5) = %d; want %d", n, 7)
	}
	if n := l.Len(); n != 5 {
		t.Errorf("Len() = %d; want %d", n, 5)
	}
	n := 50
	a := make([]int64, n)
	l = New()
	for i := range a {
		a[i] = int64(i)
		l.Append(int64(i))
	}
	if l.Len() != n {
		t.Errorf("Len() = %d; want %d", l.Len(), n)
	}
	for i := range a {
		var res int64
		for j := 0; j < i; j++ {
			res += a[j]
		}
		if l.Get(i) != a[i] {
			t.Errorf("Get(%d) = %d; want %d", i, l.Get(i), a[i])
		}
		if l.Sum(i) != res {
			t.Errorf("Sum(%d) = %d; want %d", i, l.Sum(i), res)
		}
	}

}
*/

func TestSumRange(t *testing.T) {
	l := New(1, 2, 3, -4)
	if n := l.SumRange(0, 3); n != 6 {
		t.Errorf("Range(0, 3) = %d; want %d", n, 6)
	}
	if n := l.SumRange(1, 3); n != 5 {
		t.Errorf("Range(1, 3) = %d; want %d", n, 5)
	}
	if n := l.SumRange(1, 1); n != 0 {
		t.Errorf("Range(1, 1) = %d; want %d", n, 0)
	}
	if n := l.SumRange(0, 4); n != 2 {
		t.Errorf("Range(0, 4) = %d; want %d", n, 2)
	}
	if n := l.SumRange(2, 4); n != -1 {
		t.Errorf("Range(2, 4) = %d; want %d", n, -1)
	}
}

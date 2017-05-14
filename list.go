// Package fenwick provides a list data structure with prefix sums.
//
// A Fenwick tree, or binary indexed tree, is a space-efficient list
// data structure that can efficiently update elements and calculate
// prefix sums in a list of numbers.
//
// Compared to a common array, a Fenwick tree achieves better balance
// between element update and prefix sum calculation – both operations
// run in O(log n) time – while using the same amount of memory.
// This is achieved by representing the list as an implicit tree,
// where the value of each node is the sum of the numbers in that
// subtree.
//
package fenwick

type List struct {
	data []int64
}

// New creates a new list with the given elements.
func New(n ...int64) *List {
	len := len(n)
	a := make([]int64, len)
	copy(a, n)
	for i := range a {
		if j := i | (i + 1); j < len {
			a[j] += a[i]
		}
	}
	return &List{
		data: a,
	}
}

// Len returns the number of elements in the list.
func (l *List) Len() int {
	return len(l.data)
}

// Get returns the element at index i.
func (l *List) Get(i int) int64 {
	return l.SumRange(i, i+1)
}

// Set sets the element at index i to n.
func (l *List) Set(i int, n int64) {
	l.Add(i, n-l.Get(i))
}

// Add adds n to the element at index i.
func (l *List) Add(i int, n int64) {
	for len := len(l.data); i < len; i |= i + 1 {
		l.data[i] += n
	}
}

// Sum returns the sum of the elements from index 0 to index i-1.
func (l *List) Sum(i int) int64 {
	var sum int64
	for i > 0 {
		sum += l.data[i-1]
		i -= i & -i
	}
	return sum
}

// Append appens a new element to the end of the list.
func (l *List) Append(n int64) {
	len := len(l.data)
	l.data = append(l.data, 0)
	l.Add(len, n)
}

// SumRange returns the sum of the elements from index i to index j-1.
func (l *List) SumRange(i, j int) int64 {
	var sum int64
	for j > i {
		sum += l.data[j-1]
		j -= j & -j
	}
	for i > j {
		sum -= l.data[i-1]
		i -= i & -i
	}
	return sum
}

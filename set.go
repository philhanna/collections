package collections

import (
	"cmp"
	"sort"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// Set is an unordered list of type T. Elements of type T must be Ordered
// according to the definition in https://pkg.go.dev/cmp#Ordered:
//
//	type Ordered interface {
//		~int | ~int8 | ~int16 | ~int32 | ~int64 |
//			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
//			~float32 | ~float64 |
//			~string
//	}
type Set[T cmp.Ordered] struct {
	list []T
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewSet creates a new set object of the given type
func NewSet[T cmp.Ordered](items ...T) Set[T] {
	m := make(map[T]bool)
	for _, r := range items {
		m[r] = true
	}
	this := Set[T]{}
	for k := range m {
		this.list = append(this.list, k)
	}
	return this
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Add appends an item to a set, if the set does not already contain the
// item
func (this *Set[T]) Add(item T) {
	if !this.Contains(item) {
		this.list = append(this.list, item)
	}
}

// Clear removes all elements from the set
func (this *Set[T]) Clear() {
	this.list = make([]T, 0)
}

// Contains returns true if the specified item is found in this set
func (this Set[T]) Contains(item T) bool {
	for _, v := range this.list {
		if v == item {
			return true
		}
	}
	return false
}

// Delete removes the specified item from the set. If the set does not
// contain the item, this method does nothing.
func (this *Set[T]) Delete(item T) {
	for p, thisItem := range this.list {
		if item == thisItem {
			switch p {
			case 0:
				this.list = this.list[1:]
			case len(this.list) - 1:
				this.list = this.list[:len(this.list)-1]
			default:
				prefix := this.list[:p]
				suffix := this.list[p+1:]
				this.list = append(prefix, suffix...)
			}
		}
	}
}

// Equal returns true if the two sets contain the same elements,
// regardless of order.
func (this Set[T]) Equal(that Set[T]) bool {
	return this.IsSubset(that) && that.IsSubset(this)
}

// Get return the ith element of the underlying set. This is useful
// primarily in sort functions.  Panics if i is not in the range [0 ..
// len(list)-1]
func (this Set[T]) Get(i int) T {
	return this.list[i]
}

// IsEmpty returns true if the set has no elements
func (this Set[T]) IsEmpty() bool {
	return this.Len() == 0
}

// IsSubset returns true if this set is a subset of another specified one
func (this Set[T]) IsSubset(that Set[T]) bool {
	for _, item := range this.list {
		if !that.Contains(item) {
			return false
		}
	}
	return true
}

// Iterator provides an iterator over the set, using a channel
func (this Set[T]) Iterator() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for i := 0; i < len(this.list); i++ {
			item := this.list[i]
			ch <- item
		}
	}()
	return ch
}

// IteratorSorted provides an iterator over the set, using a channel.
// The output will be ordered according to the specified less function.
func (this Set[T]) IteratorSorted(less func(i, j int) bool) <-chan T {
	// Get a sorted copy of the underlying list
	items := make([]T, 0)
	items = append(items, this.list...)
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})

	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, item := range items {
			ch <- item
		}
	}()
	return ch
}

// Len returns the number of elements in the set
func (this Set[T]) Len() int {
	return len(this.list)
}

// Map returns a new set consisting of this set with the specified
// function applied to each element
func (this Set[T]) Map(f func() T) Set[T] {
	that := NewSet[T]()
	return that
}
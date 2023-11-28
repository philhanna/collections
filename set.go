package collections

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------
type Set[T comparable] struct {
	list []T
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewSet creates a new set object of the given type
func NewSet[T comparable](items ...T) Set[T] {
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

// Len returns the number of elements in the set
func (this Set[T]) Len() int {
	return len(this.list)
}

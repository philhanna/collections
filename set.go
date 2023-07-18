package collections

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------
type Set[T comparable] []T

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
		this = append(this, k)
	}
	return this
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Contains returns true if the specified item is found in this set
func (this Set[T]) Contains(item T) bool {
	for _, v := range this {
		if v == item {
			return true
		}
	}
	return false
}

// Equal returns true if the two sets contain the same elements,
// regardless of order.
func (this Set[T]) Equal(that Set[T]) bool {
	return this.IsSubset(that) && that.IsSubset(this)
}

// IsSubset returns true if that set is a subset of this set
func (this Set[T]) IsSubset(that Set[T]) bool {
	for _, item := range that {
		if !this.Contains(item)	{
			return false
		}
	}
	return true
}

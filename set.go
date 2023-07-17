package collections

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------
type Set[T comparable] []T

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewSet creates a new set object of the given type
func NewSet[T comparable]() Set[T] {
	p := new(Set[T])
	return *p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Contains returns true if the specified item is found in this set
func (set Set[T]) Contains(item T) bool {
	for _, v := range set {
		if v == item {
			return true
		}
	}
	return false
}

// Equal returns true if the two sets contain the same elements,
// regardless of order.
func (set1 Set[T]) Equal(set2 Set[T]) bool {
	return set1.IsSubset(set2) && set2.IsSubset(set1)
}

// IsSubset returns true if set2 is a subset of set1
func (set1 Set[T]) IsSubset(set2 Set[T]) bool {
	for _, item := range set2 {
		if !set1.Contains(item)	{
			return false
		}
	}
	return true
}

// SliceToSet is a generic method that creates a set from a slice of the
// given type.
func SliceToSet[T comparable](list []T) Set[T] {
	m := make(map[T]bool)
	for _, r := range list {
		m[r] = true
	}
	set := NewSet[T]()
	for k := range m {
		set = append(set, k)
	}
	return set
}

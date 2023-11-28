package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	this := NewSet[int]()
	this.Add(3)
	this.Add(1)
	this.Add(3)
	assert.Equal(t, 2, this.Len())
}

func TestClear(t *testing.T) {
	this := NewSet[string]("Larry", "Curly", "Moe")
	this.Clear()
	assert.Equal(t, 0, this.Len())
	this.Clear()
	that := NewSet[string]()
	assert.True(t, this.Equal(that))
}

func TestContains(t *testing.T) {
	list := []string{"Larry", "Curly", "Moe", "Curly"}
	this := NewSet[string](list...)
	assert.True(t, this.Contains("Larry"))
	assert.True(t, this.Contains("Curly"))
	assert.False(t, this.Contains("Curly Joe"))
}

func TestDelete(t *testing.T) {

	getOrderedSet := func() Set[string] {
		orderedSet := NewSet[string]()
		orderedSet.Add("A")
		orderedSet.Add("B")
		orderedSet.Add("C")
		orderedSet.Add("D")
		return orderedSet
	}
	assert.Equal(t, []string{"A", "B", "C", "D"}, getOrderedSet().list)

	tests := []struct {
		name   string
		set    Set[string]
		item   string
		newSet Set[string]
	}{
		{"Delete Curly", NewSet[string]("Larry", "Curly", "Moe"), "Curly", NewSet[string]("Moe", "Larry")},
		{"Delete Curly from singleton", NewSet[string]("Curly"), "Curly", NewSet[string]()},
		{"Delete from empty set", NewSet[string](), "Bogus", NewSet[string]()},
		{"Delete first", getOrderedSet(), "A", NewSet[string]("B", "C", "D")},
		{"Delete middle", getOrderedSet(), "C", NewSet[string]("A", "B", "D")},
		{"Delete last", getOrderedSet(), "D", NewSet[string]("A", "B", "C")},
		{"Delete not contains", getOrderedSet(), "E", NewSet[string]("A", "B", "C", "D")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := tt.set
			that := tt.newSet
			this.Delete(tt.item)
			assert.True(t, this.Equal(that))
		})
	}
}

func TestIsEmpty(t *testing.T) {
	deletedSet := NewSet[int](1, 2, 3)
	deletedSet.Delete(1)
	deletedSet.Delete(2)
	deletedSet.Delete(3)

	tests := []struct {
		name string
		set  Set[int]
		want bool
	}{
		{"empty set", NewSet[int](), true},
		{"nonempty set", NewSet[int](1), false},
		{"deleted", deletedSet, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.set.IsEmpty())
		})
	}
}
func TestIsSubset(t *testing.T) {

	set1 := NewSet[string]("Larry", "Curly", "Moe")
	set2 := NewSet[string]("Larry", "Moe", "Curly")
	set3 := NewSet[string]("Larry", "Moe")
	set4 := NewSet[string]("Larry", "Jerry")
	set5 := NewSet[string]("Tom", "Dick", "Harry")
	set6 := NewSet[string]()

	tests := []struct {
		name    string
		thisSet Set[string]
		thatSet Set[string]
		want    bool
	}{
		{"Set is subset of itself", set1, set1, true},
		{"Order not important", set1, set2, true},
		{"Is proper subset", set3, set1, true},
		{"Some different elements", set4, set1, false},
		{"All different elements", set5, set1, false},
		{"Empty set is subset of any other", set6, set1, true},
		{"Empty set is subset of itself", set6, set6, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := tt.thisSet.IsSubset(tt.thatSet)
			assert.Equal(t, tt.want, have)
		})
	}
}

func TestNewSet(t *testing.T) {

	this := NewSet[string]("Larry", "Curly", "Moe", "Moe")
	assert.Equal(t, 3, len(this.list))

	that := NewSet[string]("Moe", "Larry", "Curly")
	assert.True(t, this.Equal(that))
}

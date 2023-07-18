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

func TestIsSubset(t *testing.T) {
	this := NewSet[string]("Larry", "Curly", "Moe", "Curly")
	assert.Equal(t, 3, len(this.list))
	that := NewSet[string]([]string{"Larry", "Moe", "Curly"}...)
	assert.Equal(t, 3, len(that.list))
	assert.True(t, this.IsSubset(that))
	that = NewSet[string]("Larry", "Jerry")
	assert.False(t, this.IsSubset(that))
}

func TestNewSet(t *testing.T) {

	this := NewSet[string]("Larry", "Curly", "Moe", "Moe")
	assert.Equal(t, 3, len(this.list))

	that := NewSet[string]("Moe", "Larry", "Curly")
	assert.True(t, this.Equal(that))
}

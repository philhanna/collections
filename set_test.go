package collections

import (
	"fmt"
	"strings"
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

func TestGet(t *testing.T) {
	this := NewSet[string]("B", "C", "A")
	item1 := this.Get(0)
	item2 := this.Get(1)
	item3 := this.Get(2)
	assert.True(t, this.list[0] == item1)
	assert.True(t, this.list[1] == item2)
	assert.True(t, this.list[2] == item3)
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

func TestIterator(t *testing.T) {
	this := NewSet[string]("A", "B", "C")
	that := NewSet[string]()
	for item := range this.Iterator() {
		that.Add(item)
	}
	assert.Equal(t, this, that)
}

func (this Set[T]) ExampleIterator() {
	mySet := NewSet[string]("A", "B", "C")
	for item := range mySet.Iterator() {
		fmt.Printf("%s\n", item)
	}
}

func TestMap(t *testing.T) {

	// Test with integers
	iThis := NewSet[int](3, 1, 4, 1, 5, 9)
	iHave := iThis.Map(func(n int) int {
		return n * 2
	})
	iWant := NewSet[int](6, 2, 8, 10, 18)
	assert.True(t, iHave.Equal(iWant))

	// Test with strings
	sThis := NewSet[string]("Larry", "Curly", "Moe")
	sHave := sThis.Map(func(x string) string {
		return strings.ToUpper(x)
	})
	sWant := NewSet[string]("LARRY", "CURLY", "MOE")
	assert.True(t, sHave.Equal(sWant))

	// Test with struct
	type Point struct {
		x int
		y int
	}
	pThis := NewSet[Point](Point{1, 3}, Point{2, 4})
	pHave := pThis.Map(func(p Point) Point {
		return Point{p.y, p.x}
	})
	pWant := NewSet[Point](Point{3, 1}, Point{4, 2})
	assert.True(t, pHave.Equal(pWant))

}

func TestNewSet(t *testing.T) {

	this := NewSet[string]("Larry", "Curly", "Moe", "Moe")
	assert.Equal(t, 3, len(this.list))

	that := NewSet[string]("Moe", "Larry", "Curly")
	assert.True(t, this.Equal(that))
}

func TestStructSet(t *testing.T) {
	type MyStruct struct {
		X int 
		Y int
	}
	mySet := NewSet[MyStruct]
	_ = mySet
}
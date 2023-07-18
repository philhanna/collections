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
	assert.Equal(t, 2, len(this.list))
}

func TestContains(t *testing.T) {
	list := []string{"Larry", "Curly", "Moe", "Curly"}
	set := NewSet[string](list...)
	assert.True(t, set.Contains("Larry"))
	assert.True(t, set.Contains("Curly"))
	assert.False(t, set.Contains("Curly Joe"))
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

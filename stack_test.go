package collections

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Point struct {
	Row int `json:"r"`
	Col int `json:"c"`
}

func TestStack_BadPeek(t *testing.T) {
	stack := NewStack[float64]()
	_, err := stack.Peek()
	assert.NotNil(t, err)
}

func TestStack_BadPop(t *testing.T) {
	stack := NewStack[float64]()
	_, err := stack.Pop()
	assert.NotNil(t, err)
}

func TestStack_Clear(t *testing.T) {
	stack := NewStack[float64]()
	stack.Push(math.Pi)
	stack.Push(math.E)
	assert.True(t, stack.Len() == 2)
	stack.Clear()
	assert.True(t, stack.IsEmpty())
}

func TestStack_FromJSON(t *testing.T) {
	jsonBlob := []byte(`
	[
		"Larry",
		"Curly",
		"Moe"
	]
	`)
	pst := new(Stack[string])
	pst.FromJSON(jsonBlob)

	list := pst.List
	assert.Equal(t, 3, len(list))
	assert.Equal(t, "Larry", list[0])
	assert.Equal(t, "Curly", list[1])
	assert.Equal(t, "Moe", list[2])

}

func TestStack_FromJSON_Point(t *testing.T) {
	jsonBlob := []byte(`
	[
		{"r": 1, "c": 3},
		{"r": 0, "c": 4},
		{"r": 2, "c": 1}		
	]
	`)
	pst := new(Stack[Point])
	pst.FromJSON(jsonBlob)

	list := pst.List
	assert.Equal(t, 3, len(list))
	assert.Equal(t, Point{1, 3}, list[0])
	assert.Equal(t, Point{0, 4}, list[1])
	assert.Equal(t, Point{2, 1}, list[2])

}

func TestStack_FromSlice(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("Larry")
	stack.Push("Curly")
	stack.Push("Moe")
	list := stack.List
	assert.Equal(t, 3, len(list))
	assert.Equal(t, "Larry", list[0])
	assert.Equal(t, "Curly", list[1])
	assert.Equal(t, "Moe", list[2])
}

func TestStack_GoodPeek(t *testing.T) {
	stack := NewStack[float64]()
	stack.Push(3.0)
	value, err := stack.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 3.0, value)
	assert.Equal(t, 1, stack.Len())
}

func TestStack_GoodPop(t *testing.T) {
	stack := NewStack[float64]()
	stack.Push(-17.0)
	value, err := stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, -17.0, value)
	assert.Equal(t, 0, stack.Len())
}

func TestStack_MapContentsSame(t *testing.T) {

	classic := make(map[string]int)
	classic["Larry"] = 2
	classic["Curly"] = 3
	classic["Moe"] = 1

	original := make(map[string]int)
	original["Moe"] = 1
	original["Larry"] = 2
	original["Shemp"] = 3

	noCurly := make(map[string]int)
	noCurly["Larry"] = 2
	noCurly["Moe"] = 1

	stack := NewStack[map[string]int]()
	stack.Push(original)
	stack.Push(noCurly)
	stack.Push(classic)

	assert.Equal(t, 3, stack.Len())

	var item map[string]int
	var err error

	if item, err = stack.Pop(); err != nil {
		t.Error(err)
	} else {
		// Classic
		assert.Equal(t, 1, item["Moe"])
		assert.Equal(t, 2, item["Larry"])
		assert.Equal(t, 3, item["Curly"])
	}

	if item, err = stack.Pop(); err != nil {
		t.Error(err)
	} else {
		// NoCurly
		assert.Equal(t, 1, item["Moe"])
		assert.Equal(t, 2, item["Larry"])
		_, ok := item["Curly"]
		assert.False(t, ok)
	}

	if item, err = stack.Pop(); err != nil {
		t.Error(err)
	} else {
		// Original
		assert.Equal(t, 1, item["Moe"])
		assert.Equal(t, 2, item["Larry"])
		assert.Equal(t, 3, item["Shemp"])
	}
}

func TestStack_ReverseEven(t *testing.T) {
	stack := NewStack[string]()

	stack.Push("Larry")
	stack.Push("Curly")
	stack.Push("Moe")

	stack.Reverse()

	want := []string{"Larry", "Curly", "Moe"}
	have := make([]string, 0)
	for !stack.IsEmpty() {
		s, _ := stack.Pop()
		have = append(have, s)
	}

	assert.Equal(t, want, have)
}

func TestStack_ReverseOdd(t *testing.T) {
	stack := NewStack[string]()

	stack.Push("Yesterday")
	stack.Push("Tomorrow")

	stack.Reverse()

	want := []string{"Yesterday", "Tomorrow"}
	have := make([]string, 0)
	for !stack.IsEmpty() {
		s, _ := stack.Pop()
		have = append(have, s)
	}

	assert.Equal(t, want, have)
}

func TestStack_String(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("A")
	stack.Push("B")
	assert.Equal(t, 2, stack.Len())

	var item string
	var err error

	if item, err = stack.Pop(); err != nil {
		t.Error(err)
	} else {
		letter := "B"
		assert.Equal(t, item, letter)
	}

	if item, err = stack.Pop(); err != nil {
		t.Error(err)
	} else {
		letter := "A"
		assert.Equal(t, item, letter)
	}
}

func TestStack_StructContentsSame(t *testing.T) {
	type Point struct {
		x, y int
	}
	stack := NewStack[Point]()
	stack.Push(Point{3, 4})
	if item, err := stack.Pop(); err != nil {
		t.Error(err)
	} else {
		wantX := 3
		wantY := 4
		assert.Equal(t, item.x, wantX)
		assert.Equal(t, item.y, wantY)
	}
}

func TestStack_ToJSON(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("Larry")
	stack.Push("Curly")
	stack.Push("Moe")
	jsonBlob, err := stack.ToJSON()
	assert.Nil(t, err)
	jsonstr := string(jsonBlob)
	assert.JSONEq(t, `
	[
		"Larry",
		"Curly",
		"Moe"
	]
	`, jsonstr)
}

func TestStack_ToJSON_Point(t *testing.T) {
	stack := NewStack[Point]()
	stack.Push(Point{1, 3})
	stack.Push(Point{0, 4})
	stack.Push(Point{2, 1})
	jsonBlob, err := stack.ToJSON()
	assert.Nil(t, err)
	jsonstr := string(jsonBlob)
	assert.JSONEq(t, `
	[
		{"r": 1, "c": 3},
		{"r": 0, "c": 4},
		{"r": 2, "c": 1}		
	]
	`, jsonstr)
}

func TestStack_ToSlice(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("Larry")
	stack.Push("Curly")
	stack.Push("Moe")
	list := []string{
		"Joe",
		"Smith",
	}
	stack.List = list
	assert.Equal(t, 2, stack.Len())
	entry, _ := stack.Pop()
	assert.Equal(t, "Smith", entry)
	entry, _ = stack.Pop()
	assert.Equal(t, "Joe", entry)
}

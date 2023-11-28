package collections

import (
	"encoding/json"
	"errors"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Stack is a LIFO list of items.
type Stack[T any] struct {
	List []T
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewStack creates a new stack and returns it.
func NewStack[T any]() Stack[T] {
	pst := new(Stack[T])
	return *pst
}

// ---------------------------------------------------------------------
// Error messages
// ---------------------------------------------------------------------

var ErrorStackEmpty = errors.New("Stack empty")

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Clear removes all elements from the stack.
func (pst *Stack[T]) Clear() {
	for !pst.IsEmpty() {
		pst.Pop()
	}
}

// Recreates a stack from a JSON representation.
func (pst *Stack[T]) FromJSON(jsonBlob []byte) error {
	err := json.Unmarshal(jsonBlob, &pst.List)
	return err
}

// IsEmpty returns true if the stack depth is zero.
func (pst *Stack[T]) IsEmpty() bool {
	return pst.Len() == 0
}

// Len returns the number of items in the stack.
func (pst *Stack[T]) Len() int {
	return len(pst.List)
}

// Peek returns the topmost item from the stack and returns it. The
// stack is not modified. If the stack is empty, returns an error.
func (pst *Stack[T]) Peek() (T, error) {
	empty := new(T)
	if pst.IsEmpty() {
		return *empty, ErrorStackEmpty
	}
	last := pst.Len() - 1
	result := pst.List[last]
	return result, nil
}

// Pop removes the topmost item from the stack and returns it. If the
// stack is empty, returns an error.
func (pst *Stack[T]) Pop() (T, error) {
	empty := new(T)
	if pst.IsEmpty() {
		return *empty, ErrorStackEmpty
	}
	last := pst.Len() - 1
	result := pst.List[last]
	pst.List = pst.List[:last]
	return result, nil
}

// Push appends an item onto the stack.
func (pst *Stack[T]) Push(item T) {
	pst.List = append(pst.List, item)
}

// Reverse reorders the stack contents in reverse order.
func (pst *Stack[T]) Reverse() {
	for i, j := 0, len(pst.List)-1; i < j; i, j = i+1, j-1 {
		pst.List[i], pst.List[j] = pst.List[j], pst.List[i]
	}
}

// ToJSON serializes the stack as JSON
func (pst *Stack[T]) ToJSON() ([]byte, error) {
	jsonBlob, err := json.Marshal(pst.List)
	return jsonBlob, err
}

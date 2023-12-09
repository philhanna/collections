/*
Collections is a library for using stacks and sets of any type.

Stack has the familiar operations:

	Clear() removes all elements from the stack.
	FromJSON() creates a stack from its JSON representation
	IsEmpty() returns true if the stack has no elements.
	Len() returns the number of elements in the stack.
	Peek() returns the element on the top of the stack without removing it.
	Pop() removes and returns an element from the stack,
	Push() adds an element to the stack.
	Reverse() inverts the order of the elements in the stack.
	ToJSON() creates the JSON representation of the stack

Pop() and Peek() fail when called on an empty stack.

Set is an unordered list of distinct elements. It has the following operations:

	Add() adds an element to the set, if it is not already there.
	Clear() removes all elements from the set.
	Contains() returns true if the specified item is found in this set.
	Delete() removes the specified item from the set.
	Equal() returns true if two sets contain the same elements, regardless of order.
	Filter() returns a new Set containing those elements for whom the specified function returns true.
	IsEmpty() returns true if the set has no elements.
	IsSubset() returns true if this set is a subset of another specified one.
	Iterator() iterates over each element of the underlying list
	Len() returns the number of elements in the set.
	Map() returns a new set consisting of this set with the specified function applied to each element.
*/
package collections

/*
Collections is a library for using stacks and sets of any type.

Stack has the familiar operations:

	Len() returns the number of elements in the stack.
	IsEmpty() returns true if the stack has no elements.
	Push() adds an element to the stack.
	Pop() removes and returns an element from the stack,
	Peek() returns the element on the top of the stack without removing it.
	Clear() removes all elements from the stack.

Pop() and Peek() fail when called on an empty stack.

Set is an unordered list of distinct elements. It has the following operations:

	Len() returns the number of elements in the set.
	IsEmpty() returns true if the set has no elements.
	Add() adds an element to the set, if it is not already contained there.
	Clear() removes all elements from the set.
	Contains() returns true if the specified item is found in this set.
	Delete() removes the specified item from the set.
	Equal() returns true if two sets contain the same elements, regardless of order.
	IsSubset() returns true if this set is a subset of another specified one.
*/
package collections

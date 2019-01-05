// Copyright (c) 2018 ef-ds
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package stack implements a very fast and efficient general purpose
// Last-In-First-Out (LIFO) stack data structure that is specifically
// optimized to perform when used by Microservices and serverless services
// running in production environments. This stack operates on *TestValueStack
// struct.
package stack

// TestValueStack implements an unbounded, dynamically growing Last-In-First-Out (LIFO)
// stack data structure.
// The zero value for stack is an empty stack ready to use.
type TestValueStack struct {
	// Head points to the first node of the linked list.
	head *testValueNode

	// Tail points to the last node of the linked list.
	// In an empty stack, head and tail points to the same node.
	tail *testValueNode

	// Len holds the current stack values length.
	len int
}

// TestValue is used as the value added in each push call to the queues.
// A struct is being used as structs should be more representative of real
// world uses of a queue. A second f2 field was added as the users structs
// are likely to contain more than one field.
type TestValue struct {
	F1 int
	F2 int
}

// Node represents a stack node.
// Each node holds a slice of user managed values.
type testValueNode struct {
	// v holds the list of user added values in this node.
	v []*TestValue

	// n points to the next node in the linked list.
	n *testValueNode

	// p points to the previous node in the linked list.
	p *testValueNode
}

// NewTestValueStack returns an initialized TestValueStack.
func NewTestValueStack() *TestValueStack {
	return new(TestValueStack)
}

// Init initializes or clears stack s.
func (s *TestValueStack) Init() *TestValueStack {
	*s = TestValueStack{}
	return s
}

// Len returns the number of elements of stack s.
// The complexity is O(1).
func (s *TestValueStack) Len() int { return s.len }

// Back returns the last element of stack d or nil if the stack is empty.
// The second, bool result indicates whether a valid value was returned;
// if the stack is empty, false will be returnes.
// The complexity is O(1).
func (s *TestValueStack) Back() (*TestValue, bool) {
	if s.len == 0 {
		return nil, false
	}
	return s.tail.v[len(s.tail.v)-1], true
}

// Push adds value v to the the back of the stack.
// The complexity is O(1).
func (s *TestValueStack) Push(v *TestValue) {
	switch {
	case s.head == nil:
		// No nodes present yet.
		h := &testValueNode{v: make([]*TestValue, 0, firstSliceSize)}
		h.p = h
		s.head = h
		s.tail = h
	case len(s.tail.v) < cap(s.tail.v):
		// There's room in the tail slice.
	case cap(s.tail.v) < maxInternalSliceSize:
		// We're on the first slice and it hasn't grown large enough yet.
		l := len(s.tail.v)
		nv := make([]*TestValue, l, l*sliceGrowthFactor)
		copy(nv, s.tail.v)
		s.tail.v = nv
	case s.tail.n != nil:
		// There's at least one unused slice between head and tail nodes.
		n := s.tail.n
		s.tail = n
	default:
		// No available nodes, so make one.
		n := &testValueNode{v: make([]*TestValue, 0, maxInternalSliceSize)}
		n.p = s.tail
		s.tail.n = n
		s.tail = n
	}
	s.len++
	s.tail.v = append(s.tail.v, v)
}

// Pop retrieves and removes the current element from the back of the stack.
// The second, bool result indicates whether a valid value was returned;
// if the stack is empty, false will be returnes.
// The complexity is O(1).
func (s *TestValueStack) Pop() (*TestValue, bool) {
	if s.len == 0 {
		return nil, false
	}
	s.len--
	tp := len(s.tail.v) - 1
	vp := &s.tail.v[tp]
	v := *vp
	*vp = nil // Avoid memory leaks
	s.tail.v = s.tail.v[:tp]
	if tp <= 0 {
		// Move to the previous slice as all elements
		// in the current one were removed.
		s.tail = s.tail.p
	}
	return v, true
}

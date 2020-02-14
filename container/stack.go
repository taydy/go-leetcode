package container

import "container/list"

// The <code>Stack</code> struct represents a last-in-first-out (LIFO)
// stack of objects.
type Stack struct {
	l *list.List
}

// Returns the number of objects in this stack.
func (s *Stack) Length() int {
	return s.l.Len()
}

// Removes the object at the top of this stack and returns that
// object as the value of this function.
//
// @return  The object at the top of this stack (the last item
//          of the <tt>list</tt> object).
func (s *Stack) Pop() interface{} {
	v := s.l.Back()
	s.l.Remove(v)
	return v.Value
}

// Looks at the object at the top of this stack without removing it
// from the stack.
func (s *Stack) Peek() interface{} {
	return s.l.Back().Value
}

// Pushes an item onto the top of this stack.
func (s *Stack) Push(v interface{}) {
	s.l.PushBack(v)
}

// Tests if this stack id empty.
//
// @return  <code>true</code> if and only if this stack contains
//			no items; <code>false</code> otherwise.
func (s *Stack) IsEmpty() bool {
	return s.l.Len() == 0
}

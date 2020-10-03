package stack

// Item type
type Item interface{}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	items []*Item
	count int
}

// New returns a new stack.
func (s *Stack) New() {
	s.items = make([]*Item, 0)
}

// Push adds a node to the stack.
func (s *Stack) Push(i *Item) {
	s.items = append(s.items[:s.count], i)
	s.count = s.count + 1
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Item {
	if s.count == 0 {
		return nil
	}
	length := len(s.items)
	item := s.items[length-1]

	if length > 1 {
		s.items = s.items[:length-1]
	} else {
		s.items = s.items[0:]
	}

	s.count = len(s.items)
	return item
}

// IsEmpty returns stack is empty or not
func (s *Stack) IsEmpty() bool {
	return s.count == 0
}

// Peek returns the object at the top of the Stack without removing it.
func (s *Stack) Peek() *Item {
	if s.count == 0 {
		return nil
	}

	return s.items[len(s.items)-1]
}

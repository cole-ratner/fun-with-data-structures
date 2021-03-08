package main

import "fmt"


// stack holds a slice of int
type stack struct {
	items []int
}

// push adds an item to the stack
func (s *stack) push(i int) {
	fmt.Printf("Adding: %d\n", i)
	s.items = append(s.items, i)
}

// removes the last item from the stack
func (s *stack) pop() int {
	last := len(s.items) - 1
	remove := s.items[last]
	s.items = s.items[:last]

	fmt.Printf("Removed %d from the stack\n", remove)
	return remove
}



func main() {
	s := &stack{}

	// adding items to the stack
	fmt.Println("Adding new items to the stack:")
	s.push(1)
	s.push(2)
	s.push(3)
	///

	// printing current stack
	fmt.Printf("Current stack: %v\n\n", s)

	// removing the last item from the stack
	s.pop()
	
	//printing current stack
	fmt.Printf("Current stack: %v\n", s)
}
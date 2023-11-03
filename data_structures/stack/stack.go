package main

import "fmt"

// Pushdown (Last In First Out) stack
type StackI interface {
	Push(int)      // add an item
	Pop() int      // remove the most recently added item
	IsEmpty() bool // is the stack empty?
	Size() int     // number of items in the stack
}

// The last item is on the top of the stack, so there's no need to track top index using another variable.
type IntStack []int

func (s *IntStack) Push(item int) {
	*s = append(*s, item)
}

func (s *IntStack) Pop() int {
	// * not checking for underflow
	top := (*s)[s.Size()-1]
	*s = (*s)[:s.Size()-1] // * shrink the slice
	return top
}

func (s *IntStack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *IntStack) Size() int {
	return len(*s)
}

// Create an empty stack
func Stack() StackI {
	s := IntStack([]int{})
	return &s
}

func main() {
	// create a stack
	stack := Stack()
	// test stack is empty
	if !stack.IsEmpty() {
		fmt.Printf("Stack should be empty\n")
	}

	// push items onto the stack
	items := []int{1, 2, 3}
	for _, item := range items {
		stack.Push(item)
	}

	// test stack size
	if stack.Size() != len(items) {
		fmt.Printf("Stack size should be %d\n", len(items))
	}

	// pop items off the stack
	for i := len(items) - 1; i >= 0; i-- {
		item := stack.Pop()
		if item != items[i] {
			fmt.Printf("Stack item should be %d\n", items[i])
		}
		if stack.Size() != i {
			fmt.Printf("Stack size should be %d\n", i)
		}
	}

	fmt.Println("Stack tests passed.")
}

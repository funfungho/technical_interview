package main

import "fmt"

// Fist In First Out queue
type QueueI interface {
	Enqueue(int)   // add an item
	Dequeue() int  // remove the least recently added item
	IsEmpty() bool // is the queue empty?
	Size() int     // number of items in the queue
}

// The first item will always be the head, so there's no need to track head index using another variable.
// The last item will always be the tail, so there's no need to track tail index using another variable.
type IntQueue []int

func (q *IntQueue) Enqueue(item int) {
	*q = append(*q, item)
}

func (q *IntQueue) Dequeue() int {
	// * not checking for underflow
	head := (*q)[0]
	*q = (*q)[1:] // * shrink the slice
	return head
}

func (q *IntQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *IntQueue) Size() int {
	// return len(q.Items)
	return len(*q)
}

func Queue() QueueI {
	q := IntQueue([]int{})
	return &q
}

// ! this implementation won't shrink when dequeue, memory will continue to grow
type IntQueue2 struct {
	HeadIdx int
	TailIdx int
	Items   []int
}

func main() {
	// create a queue
	queue := Queue()

	// test queue is empty
	if !queue.IsEmpty() {
		fmt.Printf("Queue should be empty\n")
	}

	// enqueue items onto the queue
	items := []int{1, 2, 3}
	for _, item := range items {
		queue.Enqueue(item)
	}

	// test queue size
	if queue.Size() != len(items) {
		fmt.Printf("Queue size should be %d\n", len(items))
	}

	// dequeue items off the queue
	for i := 0; i < len(items); i++ {
		item := queue.Dequeue()
		if item != items[i] {
			fmt.Printf("Queue item should be %d\n", items[i])
		}

		if queue.Size() != len(items)-i-1 {
			fmt.Printf("Queue size should be %d\n", len(items)-i-1)
		}
	}

	fmt.Println("Queue tests passed.")
}

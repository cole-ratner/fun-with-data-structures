package main

import "fmt"


// queue holds a slice of int
type queue struct {
	items []int
}

// enQueue adds an item to the queue
func (q *queue) enQueue(v int) {
	fmt.Printf("Now adding: %d\n", v)
	q.items = append(q.items, v)
}

// deQueue removes an item from the queue
func (q *queue) deQueue() int{
	first := q.items[0]
	q.items = q.items[1:]

	fmt.Printf("Removed: %d\n", first)
	return first
}


func main() {
	q := &queue{}

	//adding items to queue
	fmt.Println("Adding items to the queue:")
	q.enQueue(100)
	q.enQueue(200)
	q.enQueue(300)
	q.enQueue(400)
	//////

	// printing current queue state
	fmt.Printf("Current Queue: %v\n\n", q)

	// removing the first item from the queue
	q.deQueue()

	// printing current queue state
	fmt.Printf("Current Queue: %v\n\n", q)
}
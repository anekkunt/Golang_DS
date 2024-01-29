package main

import "fmt"

/*
case1: front==rear==-1 then queue empty
case2: front+1=rear then queue is full
case3: Resetting Queue (rear == front), then reset to -1
*/

const maxSize = 10

type Queue struct {
	items [maxSize]int
	front int
	rear  int
}

// NewQueue creates a new queue
func NewQueue() *Queue {
	return &Queue{front: -1, rear: -1}
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.front == -1 && q.rear == -1
}

// IsFull checks if the queue is full
func (q *Queue) IsFull() bool {
	return (q.rear+1)%maxSize == q.front
}

// Enqueue adds an item to the queue
func (q *Queue) Enqueue(item int) {
	if q.IsFull() {
		fmt.Println("Queue is full")
		return
	}
	if q.IsEmpty() {
		q.front = 0
	}
	q.rear = (q.rear + 1) % maxSize
	q.items[q.rear] = item
}

// Dequeue removes an item from the queue
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}
	item := q.items[q.front]
	if q.front == q.rear {
		// Reset queue
		q.front, q.rear = -1, -1
	} else {
		q.front = (q.front + 1) % maxSize
	}
	return item
}

func main() {
	q := NewQueue()

	// Example usage
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue()) // Outputs: 1
	fmt.Println(q.Dequeue()) // Outputs: 2
	fmt.Println(q.Dequeue()) // Outputs: Queue is empty, followed by -1
}

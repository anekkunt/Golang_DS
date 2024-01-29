package main

import "fmt"

type Queue struct {
	q []int
}

func NewQueue(cap int) *Queue {
	return &Queue{q: make([]int, 0, cap)} //make cap is 100 for efficiancy
}

func (q *Queue) enqueue(v int) {
	q.q = append(q.q, v)
}

func (q *Queue) dequeue() int {
	if len(q.q) == 0 {
		return -1
	}
	v := q.q[0]
	q.q = q.q[1:]
	return v
}

// ///////////////////////////////////////
func main() {
	myQ := NewQueue(100)
	myQ.enqueue(1)
	myQ.enqueue(2)
	myQ.enqueue(3)

	fmt.Println(myQ.dequeue())
	fmt.Println(myQ.dequeue())
	fmt.Println(myQ.dequeue())
	fmt.Println(myQ.dequeue())

}

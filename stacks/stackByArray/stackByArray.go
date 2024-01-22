package main

import (
	"errors"
	"fmt"
)

// stack is LIFO

const STACK_SIZE = 3

type stack struct {
	stack [STACK_SIZE]int
	top   int
}

func newStack() *stack {
	return &stack{
		top: -1,
	}
}
func (st *stack) Push(v int) error {
	if st.top+1 == STACK_SIZE {
		return errors.New("stack is full")
	}
	st.top++
	st.stack[st.top] = v
	return nil
}

func (st *stack) Pop() int {
	if st.top == -1 {
		return -1 //errors.New("stack is empty")
	}
	v := st.stack[st.top]
	st.top--
	return v
}

// ////////////////////////////
func main() {
	st := newStack()
	st.Push(1)
	st.Push(2)
	st.Push(3)

	fmt.Printf("%d %d %d\n", st.Pop(), st.Pop(), st.Pop())

}

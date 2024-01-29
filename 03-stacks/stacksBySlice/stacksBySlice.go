package main

import "fmt"

// stack is LIFO
type stack struct {
	stack []int
}

func newStack() *stack {
	return &stack{}
}
func (st *stack) Push(v int) {
	st.stack = append(st.stack, v)
}

func (st *stack) Pop() int {
	if len(st.stack) == 0 {
		return -1 //empty stack
	}
	top := len(st.stack) - 1
	v := st.stack[top]
	st.stack = st.stack[:top]
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

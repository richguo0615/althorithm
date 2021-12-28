package main

import "fmt"

// https://ithelp.ithome.com.tw/articles/10225040
// Golang 沒有直接的 Stack 和 Queue 可以使用，所以這邊我們自己建立了兩個 Struct：Stack 和 Queue
type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{items: make([]int, 0)}
}

func (s *Stack) push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) pop() int {
	x := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return x
}

func operateStack() {
	s := NewStack()

	times := 5

	for i := 0; i < times; i++ {
		s.push(i)
	}

	for i := 0; i < times; i++ {
		fmt.Printf("stack pop: %d \n", s.pop())
	}
}


type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{items: make([]int, 0),}
}

func (q *Queue) en(x int) {
	q.items = append(q.items, x)
}

func (q *Queue) de() int {
	if len(q.items) == 0 {
		return 0
	}

	x := q.items[0]
	q.items = q.items[1:]
	return x
}

func operateQueue() {
	q := NewQueue()

	times := 5

	for i := 0; i < times; i++ {
		q.en(i)
	}

	for i := 0; i < times; i++ {
		fmt.Printf("queue de: %d \n", q.de())
	}
}

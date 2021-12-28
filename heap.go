package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int
func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	x := (*h)[h.Len() - 1]
	*h = (*h)[:h.Len()-1]
	return x
}

func operateHeap() {
	h := &IntHeap{}
	heap.Init(h)

	times := 5

	for i := times; i >= 0; i-- {
		heap.Push(h, i)
	}

	for i := 0; i < times; i++ {
		fmt.Printf("heap pop: %d \n", heap.Pop(h).(int))
	}
}

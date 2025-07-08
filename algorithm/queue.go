package algorithm

import (
	"log"
)

type Queue struct {
	Size  int
	Num   int
	Index int
	Queue []int
}

func NewQueue(size int) *Queue {
	return &Queue{Size: size, Queue: make([]int, size)}
}

func (q *Queue) EnQueue(item int) {
	if q.Num >= q.Size {
		log.Panic("Queue overflow")
	}
	q.Queue[(q.Num+q.Index)%q.Size] = item
	q.Num += 1
}

func (q *Queue) DeQueue() int {
	if q.Num == 0 {
		log.Panic("Queue empty")
	}
	item := q.Queue[q.Index]
	q.Index = (q.Index + 1) % q.Size
	q.Num -= 1
	return item
}

func (q *Queue) GetSize() int {
	return q.Size
}

func (q *Queue) IsEmpty() bool {
	if q.Num < q.Size {
		return true
	}
	return false
}

func (q *Queue) IsFull() bool {
	return !q.IsEmpty()
}

func (q *Queue) FirstItem() int {
	return q.Queue[q.Num]
}

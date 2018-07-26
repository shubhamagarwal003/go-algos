package queue

import (
	"fmt"
)

type Queue struct {
	arr  []interface{}
	head int
	tail int
}

func (q *Queue) Push(a interface{}) {
	q.arr = append(q.arr, a)
	q.tail++
}

func (q *Queue) Pop() (interface{}, error) {
	if q.head == q.tail {
		return nil, fmt.Errorf("Queue is empty")
	}
	a := q.arr[q.head]
	q.head++
	return a, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.head == q.tail {
		return nil, fmt.Errorf("Queue is empty")
	}
	a := q.arr[q.head]
	return a, nil
}

func (q *Queue) Length() int {
	return (q.tail - q.head)
}

func NewQueue() *Queue {
	q := &Queue{head: 0, tail: 0}
	return q
}

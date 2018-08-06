package heap

import (
	"fmt"
)

type BinaryHeap struct {
	arr []int
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func (b *BinaryHeap) Print() {
	for _, i := range b.arr {
		fmt.Println(i)
	}
}

func (b *BinaryHeap) Insert(a int) {
	b.arr = append(b.arr, a)
	b.traverseUp(len(b.arr) - 1)
}

func (b *BinaryHeap) GetMin() (int, error) {
	if len(b.arr) > 0 {
		return b.arr[0], nil
	}
	return 0, fmt.Errorf("Heap is Empty")
}

func (b *BinaryHeap) ExtractMin() (int, error) {
	if len(b.arr) > 0 {
		minElement := b.arr[0]
		swap(&b.arr[0], &b.arr[len(b.arr)-1])
		b.arr = b.arr[:len(b.arr)-1]
		b.traverseDown(0)
		return minElement, nil
	}
	return 0, fmt.Errorf("Heap is Empty")
}

func (b *BinaryHeap) traverseDown(index int) {
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	length := len(b.arr)
	smallest := leftIndex
	if leftIndex < length && rightIndex < length {
		if b.arr[rightIndex] < b.arr[leftIndex] {
			smallest = rightIndex
		}
	}
	if smallest < length && b.arr[smallest] < b.arr[index] {
		swap(&b.arr[smallest], &b.arr[index])
		b.traverseDown(smallest)
	}
}

func (b *BinaryHeap) DecreaseKey(index int, amount int) {
	b.arr[index] = b.arr[index] - amount
	b.traverseUp(index)
}
func (b *BinaryHeap) traverseUp(index int) {
	parentIndex := (index - 1) / 2
	if b.arr[index] < b.arr[parentIndex] {
		swap(&b.arr[index], &b.arr[parentIndex])
		b.traverseUp(parentIndex)
	}
}

func NewHeap() *BinaryHeap {
	b := &BinaryHeap{}
	return b
}

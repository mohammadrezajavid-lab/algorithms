package algorithm

import "fmt"

type Heap struct {
	maxHeap []int
}

// GetMaxHeap for get max element of Tree by O(1)
func (heap *Heap) GetMaxHeap() []int {
	return heap.maxHeap
}
func NewMaxHeap() *Heap {
	return &Heap{maxHeap: make([]int, 0, 18)}
}
func (heap *Heap) MaxHeap() int {
	return heap.maxHeap[0]
}
func (heap *Heap) Size() int {
	return len(heap.GetMaxHeap())
}

// InsertMaxHeap = bubbleUP function is O(log n)
func (heap *Heap) InsertMaxHeap(value int) {
	heap.maxHeap = append(heap.maxHeap, value)
	valueIndex := len(heap.maxHeap) - 1
	parentIndex := (valueIndex - 1) / 2

	for valueIndex > 0 && heap.maxHeap[valueIndex] > heap.maxHeap[parentIndex] {
		heap.maxHeap[parentIndex], heap.maxHeap[valueIndex] = heap.maxHeap[valueIndex], heap.maxHeap[parentIndex]
		valueIndex = parentIndex
		parentIndex = (valueIndex - 1) / 2
	}
}

// InsertArrayBubleUp این روش در بدترین حالت O(n log n) است و در بهترین حالت زمانی که آرایه ی ورودی صعودی باشد از O(n)
func (heap *Heap) InsertArrayBubleUp(array []int) {
	for _, value := range array {
		heap.InsertMaxHeap(value)
	}
}

// InsertArray این روش در بدترین حالت O(n) است
func (heap *Heap) InsertArray(array []int) {
	heap.maxHeap = array
	arrayLength := len(array)

	for i := (arrayLength - 2) / 2; i >= 0; i-- {
		heap.bubbleDown(i)
	}
}

// DeleteMaxHeap is O(log n)
func (heap *Heap) DeleteMaxHeap() {
	if len(heap.maxHeap) == 0 {
		return
	}
	heap.maxHeap[0] = heap.maxHeap[len(heap.maxHeap)-1]
	heap.maxHeap = heap.maxHeap[0 : len(heap.maxHeap)-1]

	heap.bubbleDown(0)
}

func (heap *Heap) bubbleDown(parentIndex int) {

	leftChildIndex := (parentIndex * 2) + 1
	rightChildIndex := (parentIndex * 2) + 2
	if leftChildIndex >= heap.Size() && rightChildIndex >= heap.Size() {
		return
	}

	maxChildIndex := heap.maxChild(leftChildIndex, rightChildIndex)

	if heap.maxHeap[parentIndex] >= heap.maxHeap[maxChildIndex] {
		return
	}

	heap.maxHeap[parentIndex], heap.maxHeap[maxChildIndex] = heap.maxHeap[maxChildIndex], heap.maxHeap[parentIndex]
	heap.bubbleDown(maxChildIndex)
}

func (heap *Heap) maxChild(leftChildIndex, rightChildIndex int) int {
	if leftChildIndex < heap.Size() && rightChildIndex >= heap.Size() {
		if heap.maxHeap[leftChildIndex] > heap.maxHeap[(leftChildIndex-1)/2] {
			return leftChildIndex
		} else {
			return (leftChildIndex - 1) / 2
		}
	}
	if heap.maxHeap[leftChildIndex] > heap.maxHeap[rightChildIndex] {
		return leftChildIndex
	}
	return rightChildIndex
}
func (heap *Heap) PrintMaxHeap() {
	for _, value := range heap.maxHeap {
		fmt.Printf("%d ", value)
	}
}

package algorithm

import (
	"math"
	"math/rand/v2"
	"slices"
)

type QuickSort struct {
	array []int
}

func NewQuickSort(intPutArray []int) *QuickSort {
	return &QuickSort{array: intPutArray}
}
func (quick *QuickSort) SetArray(array []int) {
	quick.array = array
}
func (quick *QuickSort) GetArray() []int {
	return quick.array
}
func (quick *QuickSort) swap(index1, index2 int) {
	quick.array[index1], quick.array[index2] = quick.array[index2], quick.array[index1]
}

// Sort is quickSort method by O(n log n)
func (quick *QuickSort) Sort(left, right int) {
	if right <= left {
		return
	}
	pivotIndex := quick.partition(left, right)
	quick.Sort(left, pivotIndex-1)
	quick.Sort(pivotIndex+1, right)
}
func (quick *QuickSort) partition(left, right int) int {
	randomPivotIndex := func(minNum, maxNum int) int {
		return rand.IntN(maxNum-minNum) + minNum
	}

	quick.swap(randomPivotIndex(left, right), right)
	pivot := quick.array[right]
	leftP := left
	rightP := right - 1

	for leftP <= rightP {
		for leftP <= rightP && quick.array[leftP] <= pivot {
			leftP += 1
		}
		for leftP <= rightP && quick.array[rightP] > pivot {
			rightP -= 1
		}
		if leftP < rightP {
			quick.swap(leftP, rightP)
		}
	}
	quick.swap(leftP, right)
	return leftP
}

// EfficientSort
func (quick *QuickSort) EfficientSort(left, right int) {
	if right <= left {
		return
	}
	pivotIndex := quick.efficientPartition(left, right)
	quick.EfficientSort(left, pivotIndex-1)
	quick.EfficientSort(pivotIndex+1, right)
}

// efficient Partition
func (quick *QuickSort) efficientPartition(left, right int) int {
	pivotIndex := selectEfficientPivot(quick.array)
	quick.swap(pivotIndex, right)
	pivot := quick.array[right]
	leftP := left
	rightP := right - 1

	for leftP <= rightP {
		for leftP <= rightP && quick.array[leftP] <= pivot {
			leftP += 1
		}
		for leftP <= rightP && quick.array[rightP] > pivot {
			rightP -= 1
		}
		if leftP < rightP {
			quick.swap(leftP, rightP)
		}
	}
	quick.swap(leftP, right)
	return leftP
}
func selectEfficientPivot(array []int) int {
	var numberOfSets int = len(array) / 5
	start, end := 0, 5
	tmpArray := array
	for len(tmpArray) > 5 {
		tmp := make([]int, 0, 10)
		for i := 1; i <= numberOfSets; i++ {
			tmp = append(tmp, mergeSort(array[start:end])[2])
			start = end
			end = start + 5
		}
		tmpArray = tmp
	}
	return slices.Index(array, tmpArray[len(tmpArray)/2]) - 1
}

// Get این متد k امین عنصر آرایه را به ما میدهد در صورتی که آرایه ی ما نامرتب باشد
func (quick *QuickSort) Get(k int) int {
	return quick.get(k, 0, len(quick.GetArray())-1)
}
func (quick *QuickSort) get(k, left, right int) int {
	if right <= left {
		return math.MinInt
	}
	pivotIndex := quick.partition(left, right)
	switch {
	case k == pivotIndex:
		return quick.GetArray()[pivotIndex]
	case k < pivotIndex:
		return quick.get(k, left, pivotIndex-1)
	case k > pivotIndex:
		return quick.get(k, pivotIndex+1, right)
	default:
		return math.MinInt
	}
}

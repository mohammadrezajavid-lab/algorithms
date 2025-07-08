package algorithm

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
)

// 2T(n/2) -> O(log n)
func binarySearch(inputArray []int, value int, start int, end int) int {
	var midIndex = (start + end) / 2

	if start > end {
		return -1
	}
	if inputArray[start] == value {
		return start
	}
	if inputArray[end] == value {
		return end
	}
	if inputArray[midIndex] == value {
		return midIndex
	}

	if value < inputArray[midIndex] {
		return binarySearch(inputArray, value, start, midIndex-1)
	}
	if value > inputArray[midIndex] {
		return binarySearch(inputArray, value, midIndex+1, end)
	}

	return -1
}

func LowerBoundBinarySearch(nums []int, x int) (int, int) {
	if x <= nums[0] {
		return -1, -1
	}
	start := 0
	end := len(nums)
	for start+1 < end {
		mid := (start + end) >> 1
		if x <= nums[mid] {
			end = mid
		} else {
			start = mid
		}
	}
	return start, nums[start]
}

func UpperBoundBinarySearch(nums []int, x int) (int, int) {
	if x >= nums[len(nums)-1] {
		return -1, -1
	}
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := (start + end) >> 1
		if x < nums[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start, nums[start]
}

func MagicFunctionBinarySearch(nums []int, x int) (int, int) {
	if x >= nums[0] {

		return -1, -1
	}
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := (start + end + 1) >> 1
		if x < nums[mid] {
			start = mid
		} else {
			end = mid - 1
		}
	}
	return start, nums[start]
}

func BinarySearchInsert(nums []int, target int) int {

	var start int = 0
	var end int = len(nums) - 1
	var midIndex int = -1
	var outputIndex int = -1

	for start <= end {
		midIndex = (start + end) / 2
		if start == end {
			if target <= nums[midIndex] {
				outputIndex = midIndex
				break
			} else if target > nums[midIndex] {
				outputIndex = midIndex + 1
				break
			}
		}
		if nums[midIndex] == target {
			outputIndex = midIndex
			break
		}
		if target < nums[midIndex] {
			end = midIndex
			continue
		}
		if target > nums[midIndex] {
			start = midIndex + 1
		}

	}
	return outputIndex
}

func bucketSort(inputArray []int) []int {
	buckets := make([][]int, 10)

	for _, value := range inputArray {
		buckets[value/10] = append(buckets[value/10], value)
	}

	inputArray = nil
	for _, bucket := range buckets {
		inputArray = append(inputArray, bubbleSort(bucket)...)
	}

	return inputArray
}

/*
* T(n):	O(1)	if n == 1
* T(n):	2T(n/2) + O(n)	if n > 1

* 		Time complexity
* best-case: 	O(n log n) 	-> if the list is already sorted
* worst-case: 	O(n log n) 	-> if the list is in reverse ordered
* average-case: O(n log n)  -> if the list is randomly ordered

* 		Space complexity
* Auxiliary space: O(n) 	-> Additional space for temporary array use during merging.
 */

func mergeSort(inputArray []int) []int {
	if len(inputArray) <= 2 {
		if len(inputArray) == 2 {
			if inputArray[1] < inputArray[0] {
				inputArray[0], inputArray[1] = inputArray[1], inputArray[0]
			}
		}
		return inputArray
	}

	// Divide
	mid := len(inputArray) / 2
	divideL := mergeSort(inputArray[:mid]) // T(n/2) -> each half has n/2 elements,
	// we have two recursive calls with input size as (n/2).
	divideR := mergeSort(inputArray[mid:]) // T(n/2)

	var result []int
	var i, j = 0, 0

	// Conquer and Merge // O(n) -> for merge the two sorted halves
	for i < len(divideL) && j < len(divideR) {
		if divideL[i] < divideR[j] {
			result = append(result, divideL[i])
			i += 1
			continue
		}
		if divideR[j] < divideL[i] {
			result = append(result, divideR[j])
			j += 1
			continue
		}
		if divideL[i] == divideR[j] {
			result = append(result, divideL[i])
			i += 1
			continue
		}
	}
	if i < len(divideL) {
		result = append(result, divideL[i:]...)
	}
	if j < len(divideR) {
		result = append(result, divideR[j:]...)
	}

	return result
}

/*
* 		Time complexity
* best-case: 	O(n^2) 	-> if the list is already sorted
* worst-case: 	O(n^2) 	-> if the list is in reverse ordered
* average-case: O(n^2)  -> if the list is randomly ordered

* 		Space complexity
* Auxiliary space: O(1)
 */
func selectionSort(inputArray []int) []int {
	for i := 0; i < len(inputArray)-1; i++ {
		minItem := slices.Min(inputArray[i+1:])
		minItemIndex := slices.Index(inputArray, minItem)
		if minItem < inputArray[i] {
			inputArray[i], inputArray[minItemIndex] = inputArray[minItemIndex], inputArray[i]
		}
	}
	return inputArray
}

// Bubble sort by swapping
func bubbleSort(inputArray []int) []int {
	for i := 0; i < len(inputArray); i++ {
		var bubbleFound = false
		for j := len(inputArray) - 1; j > i; j-- {
			if inputArray[j] < inputArray[j-1] {
				inputArray[j], inputArray[j-1] = inputArray[j-1], inputArray[j]
				bubbleFound = true
			}
		}
		if !bubbleFound {
			break
		}
	}
	return inputArray
}

// Bubble sort by shifting
/*
* 		Time complexity
* best-case: 	O(n) 	-> if the list is already sorted
* worst-case: 	O(n^2) 	-> if the list is in reverse ordered
* average-case: O(n^2)  -> if the list is randomly ordered

* 		Space complexity
* Auxiliary space: O(1)
 */
func bubbleSortShift(inputArray []int) []int {
	for i := 0; i < len(inputArray); i++ {
		minItem := inputArray[len(inputArray)-1]
		var j = len(inputArray) - 1
		var bubbleFound = false
		for ; j > i; j-- {
			if minItem < inputArray[j-1] {
				inputArray[j] = inputArray[j-1]
				bubbleFound = true
			} else if minItem > inputArray[j-1] {
				inputArray[j] = minItem
				minItem = inputArray[j-1]
			}
		}
		inputArray[j] = minItem
		if !bubbleFound {
			break
		}
	}
	return inputArray
}

/*
* 		Time complexity
* best-case: 	O(n) 	-> if the list is already sorted
* worst-case: 	O(n^2) 	-> if the list is in reverse ordered
* average-case: O(n^2)  -> if the list is randomly ordered

* 		Space complexity
* Auxiliary space: O(1)
 */
func bestInsertionSort(inputArray []int) []int {
	for i := 1; i < len(inputArray); i++ {
		item := inputArray[i]
		var j = i - 1
		for ; j >= 0; j-- {
			if item < inputArray[j] {
				inputArray[j+1] = inputArray[j]
			} else {
				break
			}
		}
		inputArray[j+1] = item
	}
	return inputArray
}

func insertionSort(inputArray []int) []int {
	for i := 1; i < len(inputArray); i++ {
		for j := i; j > 0; j-- {
			if inputArray[j] < inputArray[j-1] {
				inputArray[j-1], inputArray[j] = inputArray[j], inputArray[j-1]
			}
		}

	}
	return inputArray
}

/*
* N -> size of inputArray[]
* M -> size of countArray[]
* and Auxiliary Space: O(N+M) N and M for countArray, sortedArray
* Worst-case: O(N+M)
* Average-case: O(N+M)
* Best-case: O(N+M)
 */
func countSort(inputArray []int) []int {
	maxItem := slices.Max(inputArray)
	countArray := make([]int, maxItem+1)

	for _, item := range inputArray {

		countArray[item]++
	}

	for i := 1; i < len(countArray); i++ {

		countArray[i] += countArray[i-1]
	}

	sortedArray := make([]int, len(inputArray))
	for i := len(inputArray) - 1; i >= 0; i-- {
		countArray[inputArray[i]] -= 1
		sortedArray[countArray[inputArray[i]]] = inputArray[i]
	}

	return sortedArray
}

func countingSortFor(inputArray []int) []int {
	maxItem := slices.Max(inputArray)
	countArray := make([]int, maxItem+1)

	for _, item := range inputArray {
		countArray[item]++
	}

	sortedArray := make([]int, len(inputArray))
	var k = 0
	for i := 0; i < len(countArray); i++ {
		for j := 0; j < countArray[i]; j++ {
			sortedArray[k], k = i, k+1
		}
	}

	return sortedArray
}

/*
این روش از مرتب سازی کاربردش برای مقادیر گسسته است
منظور اینه که چند تا مقادیر مشخص داشته باشیم که بشه بین دسته های مشخص تقسیمشون کنیم
تعداد دسته هامون محدود باشه مثلا اگه تعداد دستهای ما 5 ملیارد باشه ما باید یک آرایه ی 5 ملیاردی در نظر بگیریم که این خیلی بده
مثلا: اگه به ما 5 عدد داده باشن که مقادیرشون بین 1 تا 100 ملیارد هست ایا این میصرفه که از مرتب سازی شمارشی استفاده کنیم؟ نه
پس باید 2تا شرط زیر رو داشته باشیم که بشه از روش مرتب سازی شمارشی استفاده کرد:
1- ورودی ها گسسته و قابل شمارش باشن
2- اگر تعداد عنصرهای ورودی n باشن و تعداد دسته ها m باشن شرط زیر باید برقرار باشه
n >> m

وردی ما: فقط آرایه ای از عددهای مثبت
*/
// Counting sort takes O(n+k) time and O(n+k) space
// n is the number of items we're sorting
// k is the number of possible values
func countingSort(inputArray []uint) {

	maxItem := slices.Max(inputArray)
	count := make([]uint, maxItem+1)

	for _, u := range inputArray {
		count[u]++
	}

	var builder strings.Builder
	for i, measure := range count {
		stringNumber := strings.Repeat(strconv.Itoa(i)+" ", int(measure))
		builder.WriteString(stringNumber)
	}

	fmt.Println(builder.String())
}

func generateRandomArray(size, maxValue int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(maxValue)
	}
	return arr
}

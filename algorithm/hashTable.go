package algorithm

import (
	"log"
	"math"
)

type (
	HashTableNode struct {
		status bool // status, if one node is nil, status is true, and if not nil is false.
		data   int
		list   *LinkedList
	}
	HashTable struct {
		array []*HashTableNode
		count int
	}
)

func NewHashTable() *HashTable {
	return &HashTable{
		array: make([]*HashTableNode, 101),
	}
}
func NewHashTableNode(data int) *HashTableNode {
	return &HashTableNode{
		status: false,
		data:   data,
		list:   NewLinkedList(),
	}
}

func (htn *HashTableNode) GetData() int {
	return htn.data
}
func (htn *HashTableNode) GetStatus() bool {
	return htn.status
}
func (htn *HashTableNode) SetStatus(status bool) {
	htn.status = status
}
func (htn *HashTableNode) GetList() *LinkedList {
	return htn.list
}
func (htn *HashTableNode) SetData(data int) {
	htn.data = data
}
func (hash *HashTable) incrementCount() {
	hash.count += 1
}
func (hash *HashTable) decrementCount() {
	hash.count -= 1
}
func (hash *HashTable) GetArray() []*HashTableNode {
	return hash.array
}
func (hash *HashTable) setArray(array []*HashTableNode) {
	hash.array = array
}
func (hash *HashTable) GetSize() int {
	return len(hash.array)
}
func (hash *HashTable) GetCount() int {
	return hash.count
}
func (hash *HashTable) Insert(data int) {
	if hash.loadFactor() {
		hash.resize()
	}
	var index int = hash.hashModulo(data)
	if hash.array[index] == nil {
		hash.array[index] = NewHashTableNode(data)
		hash.incrementCount()
		return
	}
	if hash.array[index].GetStatus() {
		hash.array[index].SetData(data)
		hash.incrementCount()
	} else {
		hash.array[index].list.Append(data)
		hash.incrementCount()
	}
}
func (hash *HashTable) Find(data int) bool {
	var index int = hash.hashModulo(data)
	if hash.array[index] == nil {
		return false
	} else {
		if hash.array[index].GetData() == data {
			return true
		} else {
			if hash.array[index].GetList().Index(data) == -1 {
				return false
			}
		}
	}
	return true
}
func (hash *HashTable) Delete(data int) {
	var index int = hash.hashModulo(data)
	if hash.array[index] == nil {
		log.Fatalf("This data[%v] not yet.", data)
	} else {
		if hash.array[index].GetData() == data {
			hash.array[index].SetStatus(true)
			hash.decrementCount()
		} else {
			hash.array[index].GetList().Delete(data)
			hash.decrementCount()
		}
	}
}
func (hash *HashTable) hashModulo(key int) int {
	return key % hash.GetSize()
}
func (hash *HashTable) multiplicationHashing(key int) int {
	A := 0.6180339887
	return int(A*float64(key)) % hash.GetSize()
}
func (hash *HashTable) divisionHashing(key int) int {
	return int(float64(hash.GetSize()) * float64(float64(hash.GetSize())/float64(key)))
}

// FoldingHashing کاربردش برای کلید های طولانی مثل کد ملی یا شماره تلفن خوبه
func (hash *HashTable) foldingHashing(key int) int {
	// 12345678 --> (1234 + 5678) % M :::: 12345678 - (12345678/10000)*10000
	left := key / 10000
	right := key - (key/10000)*10000

	return (left + right) % hash.GetSize()
}

func (hash *HashTable) resize() {
	var newSize int = nextPrime(hash.GetSize())
	var oldArray []*HashTableNode = hash.GetArray()
	var oldArraySize int = hash.GetSize()
	hash.setArray(make([]*HashTableNode, newSize))
	for i := 0; i < oldArraySize; i++ {
		if oldArray[i] != nil {
			if !oldArray[i].status {
				hash.Insert(oldArray[i].data)
			}
			if oldArray[i].list.Size() > 0 {
				for j := 0; j < oldArray[i].GetList().Size(); j++ {
					hash.Insert(oldArray[i].GetList().Pop())
				}
			}
		}
	}
}

// loadFactor : if load factor value > 2 -> hashTable must resize
func (hash *HashTable) loadFactor() bool {
	if hash.GetCount()/hash.GetSize() > 2 {
		return true
	}
	return false
}

// NextPrime اول عدد رو دوبرابر میکنه بعد اولین عدد اول بعدش رو میده
func nextPrime(num int) int {
	num *= 2
	for {
		num++
		if isPrime(num) {
			return num
		}
	}
}
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

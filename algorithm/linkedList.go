package algorithm

import (
	"log"
)

type (
	LinkedListNode struct {
		pre  *LinkedListNode
		next *LinkedListNode
		data int
	}
	LinkedList struct {
		head *LinkedListNode
		tail *LinkedListNode
		size int
	}
)

func NewLinkedListNode(data int) *LinkedListNode {
	return &LinkedListNode{data: data}
}
func NewLinkedList() *LinkedList {
	tail := &LinkedListNode{pre: nil, next: nil, data: 0}
	head := &LinkedListNode{pre: nil, next: tail, data: 0}
	tail.pre = head
	return &LinkedList{head: head, tail: tail}
}
func (linkedList *LinkedList) Size() int {
	return linkedList.size
}
func (linkedList *LinkedList) Head() *LinkedListNode {
	return linkedList.head
}
func (linkedList *LinkedList) Tail() *LinkedListNode {
	return linkedList.tail
}
func (linkedList *LinkedList) Append(data int) {
	var newNode *LinkedListNode = NewLinkedListNode(data)
	newNode.next = linkedList.tail
	newNode.pre = linkedList.tail.pre
	linkedList.tail.pre.next = newNode
	linkedList.tail.pre = newNode
	linkedList.size += 1
}
func (linkedList *LinkedList) Insert(data, index int) {
	var newNode *LinkedListNode = NewLinkedListNode(data)
	if index < 0 || index > linkedList.Size() {
		log.Panicf("Index[%d] out of range", index)
		return
	}
	for i := 0; i < linkedList.Size(); i++ {
		if i == index {
			newNode.next = linkedList.tail
			newNode.pre = linkedList.tail.pre
			linkedList.tail.pre.next = newNode
			linkedList.tail.pre = newNode
			linkedList.size += 1
		}
		return
	}
}
func (linkedList *LinkedList) Delete(data int) {
	var node *LinkedListNode
	for i := 0; i < linkedList.size; i++ {
		if data == node.data {
			break
		}
		node = node.next
	}
	node.pre.next = node.next
	node.next.pre = node.pre
	linkedList.size -= 1
}
func (linkedList *LinkedList) Pop() int {
	dataDeleted := linkedList.Tail().pre.data
	linkedList.Tail().pre.pre.next = linkedList.Tail()
	linkedList.Tail().pre = linkedList.Tail().pre.pre
	return dataDeleted
}
func (linkedList *LinkedList) Index(data int) int {
	var tempNode *LinkedListNode = linkedList.head.next
	for i := 0; i < linkedList.size; i++ {
		if tempNode.data == data {
			return i
		}
		tempNode = tempNode.next
	}
	return -1
}
func (linkedList *LinkedList) Get(index int) *LinkedListNode {
	var tempNode *LinkedListNode = linkedList.head.next
	for i := 0; i < linkedList.size; i++ {
		if i == index {
			return tempNode
		}
		tempNode = tempNode.next
	}
	return nil
}

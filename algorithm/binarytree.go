package algorithm

import (
	"fmt"
	"log"
)

type Node struct {
	Data       int
	Parent     *Node
	LeftChild  *Node
	RightChild *Node
}
type BinaryTree struct {
	Root *Node
	Size int // number of nodes in binary tree
}

func (bTree *BinaryTree) SetRoot(rootNode *Node) {
	bTree.Root = rootNode
}

func (bTree *BinaryTree) GetRoot() *Node {
	return bTree.Root
}
func (n *Node) Print() {
	fmt.Printf("Data: %d\t Parent: %v\t LeftChild: %v\t RightChild: %v\n", n.Data, n.Parent, n.LeftChild, n.RightChild)
}

func NewNode(data int, parent *Node) *Node {
	return &Node{Data: data, Parent: parent, LeftChild: nil, RightChild: nil}
}

func NewBinaryTree(rootNode *Node) *BinaryTree {
	return &BinaryTree{Root: rootNode, Size: 1}
}

func (bTree *BinaryTree) Insert(data int, rootNode *Node) {
	if data <= rootNode.Data {
		if rootNode.LeftChild == nil {
			rootNode.LeftChild = NewNode(data, rootNode)
			bTree.Size += 1
			return
		}
		bTree.Insert(data, rootNode.LeftChild)
	} else {
		if rootNode.RightChild == nil {
			rootNode.RightChild = NewNode(data, rootNode)
			bTree.Size += 1
			return
		}
		bTree.Insert(data, rootNode.RightChild)
	}
}

func (bTree *BinaryTree) Search1(data int, rootNode *Node) bool {
	if rootNode == nil {
		return false
	}
	if data == rootNode.Data {
		return true
	}
	return bTree.Search1(data, rootNode.LeftChild) || bTree.Search1(data, rootNode.RightChild)
}

func (bTree *BinaryTree) Search2(data int, rootNode *Node) *Node {
	if rootNode == nil {
		return nil
	}
	if rootNode.Data == data {
		return rootNode
	}
	if data < rootNode.Data {
		return bTree.Search2(data, rootNode.LeftChild)
	}
	return bTree.Search2(data, rootNode.RightChild)
}

func (bTree *BinaryTree) InOrder(node *Node) {
	if node == nil {
		return
	}
	bTree.InOrder(node.LeftChild)
	fmt.Printf("%d ", node.Data)
	bTree.InOrder(node.RightChild)
}

func (bTree *BinaryTree) PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Data)
	bTree.PreOrder(node.LeftChild)
	bTree.PreOrder(node.RightChild)
}

func (bTree *BinaryTree) PostOrder(node *Node) {
	if node == nil {
		return
	}
	bTree.PostOrder(node.LeftChild)
	bTree.PostOrder(node.RightChild)
	fmt.Printf("%d ", node.Data)
}

func (bTree *BinaryTree) Remove(data int, rootNode *Node) {
	findNode := bTree.Search2(data, rootNode)

	bTree.removeNode(findNode)
}

func (bTree *BinaryTree) removeNode(n *Node) {
	switch {
	case n.LeftChild == nil && n.RightChild == nil:
		if n.Parent.LeftChild != nil {
			if n.Parent.LeftChild.Data == n.Data {
				n.Parent.LeftChild = nil
			}
		} else if n.Parent.RightChild != nil {
			if n.Parent.RightChild.Data == n.Data {
				n.Parent.RightChild = nil
			}
		} else {
			log.Panicf("Not yet this Node parent: %d", n.Data)
		}
	case n.LeftChild != nil && n.RightChild == nil:
		if n.Parent.LeftChild != nil && n.Parent.LeftChild.Data == n.Data {
			n.LeftChild.Parent = n.Parent
			n.Parent.LeftChild = n.LeftChild
		} else if n.Parent.RightChild != nil && n.Parent.RightChild.Data == n.Data {
			n.LeftChild.Parent = n.Parent
			n.Parent.RightChild = n.LeftChild
		}
	case n.LeftChild == nil && n.RightChild != nil:
		if n.Parent.LeftChild != nil && n.Parent.LeftChild.Data == n.Data {
			n.RightChild.Parent = n.Parent
			n.Parent.LeftChild = n.RightChild
		} else if n.Parent.RightChild != nil && n.Parent.RightChild.Data == n.Data {
			n.RightChild.Parent = n.Parent
			n.Parent.RightChild = n.RightChild
		}
	case n.LeftChild != nil && n.RightChild != nil:
		selectedNode := bTree.selectNode(n.LeftChild)

		n.Data, selectedNode.Data = selectedNode.Data, n.Data
		bTree.removeNode(selectedNode)
	default:
		log.Printf("Node with data %d not found!", n.Data)
		return
	}
}

// SelectNode This function selected max node in the left child tree; this function is O(log n)
func (bTree *BinaryTree) selectNode(node *Node) *Node {
	if node == nil {
		return nil
	}
	for node.RightChild != nil {
		node = node.RightChild
	}
	return node
}

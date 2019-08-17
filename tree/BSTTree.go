package tree

import (
	"fmt"
)

type BSTTree struct {
	root *Node
}

func (t *BSTTree) Insert(e interface{}) {
	if t.root == nil {
		t.root = &Node{}
		t.root.Value = e
	} else {
		objectValue := t.root.Value.(int)
		if objectValue > e.(int) {
			if t.root.Left == nil {
				newNode := Node{}
				newNode.Parent = t.root
				newNode.Value = e
				t.root.Left = &newNode
			} else {
				t.insertToNode(t.root.Left, e)
			}
		} else {
			if t.root.Right == nil {
				newNode := Node{}
				newNode.Parent = t.root
				newNode.Value = e
				t.root.Right = &newNode
			} else {
				t.insertToNode(t.root.Right, e)
			}
		}
	}
}

func (t *BSTTree) insertToNode(node *Node, e interface{}) {
	objectValue := node.Value.(int)
	if objectValue > e.(int) {
		if node.Left != nil {
			t.insertToNode(node.Left, e)
		} else {
			newNode := Node{}
			newNode.Parent = node
			newNode.Value = e
			node.Left = &newNode
		}
	} else {
		if node.Right != nil {
			t.insertToNode(node.Right, e)
		} else {
			newNode := Node{}
			newNode.Parent = node
			newNode.Value = e
			node.Right = &newNode
		}
	}

}

func (t *BSTTree) PreOrderPrint() {
	if t.root.Left != nil {
		t.preOrderPrintFromNode(t.root.Left)
	}
	fmt.Println(t.root.Value)
	if t.root.Right != nil {
		t.preOrderPrintFromNode(t.root.Right)
	}
}

func (t *BSTTree) preOrderPrintFromNode(node *Node) {
	if node.Left != nil {
		t.preOrderPrintFromNode(node.Left)
	}
	fmt.Println(node.Value)
	if node.Right != nil {
		t.preOrderPrintFromNode(node.Right)
	}
}

func (t *BSTTree) Search(e interface{}) *Node {
	if e.(int) == t.root.Value.(int) {
		return t.root
	} else if e.(int) < t.root.Value.(int) {
		return t.searchInSubtree(e, t.root.Left)
	} else {
		return t.searchInSubtree(e, t.root.Right)
	}
}

func (t *BSTTree) searchInSubtree(e interface{}, root *Node) *Node {
	if root == nil {
		return nil
	}
	if e.(int) == root.Value.(int) {
		return root
	} else if e.(int) < root.Value.(int) {
		return t.searchInSubtree(e, root.Left)
	} else {
		return t.searchInSubtree(e, root.Right)
	}
}

func (t *BSTTree) findInOrderSuccessor(node *Node) *Node {
	return t.findMinimumSubTree(node)
}

func (t *BSTTree) findMinimumSubTree(node *Node) *Node {
	if node.Left != nil {
		return t.findMinimumSubTree(node.Left)
	}
	return node
}

func (t *BSTTree) Delete(e *Node) {
	if e.Left == nil && e.Right == nil {
		e = nil
	} else if e.Left != nil && e.Right == nil {
		parent := e.Parent
		if parent.Left == e {
			parent.Left = e.Left
		} else {
			parent.Right = e.Left
		}
		e = nil
	} else if e.Left == nil && e.Right != nil {
		parent := e.Parent
		if parent.Left == e {
			parent.Left = e.Right
		} else {
			parent.Right = e.Right
		}
		e = nil
	} else {
		successor := t.findInOrderSuccessor(e.Right)
		e.Value = successor.Value
		parent := successor.Parent
		if parent.Left == successor {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	}
}

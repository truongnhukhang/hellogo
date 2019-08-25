package tree

import (
	"fmt"
	"github.com/truongnhukhang/hellogo/queue"
	"strconv"
)

type RBTree struct {
	root *Node
}

func (t *RBTree) Insert(e interface{}) {
	if t.root == nil {
		t.root = &Node{}
		t.root.Value = e
		t.root.Color = "b"
	} else {
		objectValue := t.root.Value.(int)
		if objectValue > e.(int) {
			if t.root.Left == nil {
				newNode := Node{}
				newNode.Parent = t.root
				newNode.Value = e
				newNode.Color = "r"
				t.root.Left = &newNode
			} else {
				t.insertFixup(t.insertToNode(t.root.Left, e))
			}
		} else {
			if t.root.Right == nil {
				newNode := Node{}
				newNode.Parent = t.root
				newNode.Value = e
				newNode.Color = "r"
				t.root.Right = &newNode
			} else {
				t.insertFixup(t.insertToNode(t.root.Right, e))
			}
		}

	}

}

func (t *RBTree) insertFixup(node *Node) {
	for node != t.root && node.Parent.Color == "r" {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == "r" {
				node.Parent.Color = "b"
				uncle.Color = "b"
				node.Parent.Parent.Color = "r"
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					node = node.Parent
					t.leftRotate(node)
				}
				node.Parent.Color = "b"
				node.Parent.Parent.Color = "r"
				t.rightRotate(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == "r" {
				node.Parent.Color = "b"
				uncle.Color = "b"
				node.Parent.Parent.Color = "r"
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					t.rightRotate(node)
				}
				node.Parent.Color = "b"
				node.Parent.Parent.Color = "r"
				t.leftRotate(node.Parent.Parent)
			}
		}
		t.root.Color = "b"
	}
}

func (t *RBTree) leftRotate(node *Node) {
	if node.Right != nil {
		right := node.Right
		node.Right = right.Left
		right.Left = node
		parent := node.Parent
		if parent == nil {
			t.root = right
			right.Parent = nil
		} else {
			if parent.Left == node {
				parent.Left = right
			} else {
				parent.Right = right
			}
			node.Parent = right
			right.Parent = parent
		}
	}
}

func (t *RBTree) rightRotate(node *Node) {
	if node.Left != nil {
		left := node.Left
		node.Left = left.Right
		left.Right = node
		parent := node.Parent
		if parent == nil {
			t.root = left
			left.Parent = nil
		} else {
			if parent.Left == node {
				parent.Left = left
			} else {
				parent.Right = left
			}
			node.Parent = left
			left.Parent = parent
		}
	}
}

func (t *RBTree) LevelOrderPrint() {
	fmt.Println("************************")
	simpleQueue := queue.SimpleQueue{[]interface{}{}}
	simpleQueue.Put(t.root)
	var tempNode *Node = nil
	for !simpleQueue.IsEmpty() {
		tempNode = simpleQueue.Poll().(*Node)
		fmt.Println(strconv.Itoa(tempNode.Value.(int)) + " - " + tempNode.Color)
		if tempNode.Left != nil {
			simpleQueue.Put(tempNode.Left)
		}
		if tempNode.Right != nil {
			simpleQueue.Put(tempNode.Right)
		}
	}
}

func (t *RBTree) PreOrderPrint() {
	if t.root.Left != nil {
		t.preOrderPrintFromNode(t.root.Left)
	}
	fmt.Println(t.root.Value)
	if t.root.Right != nil {
		t.preOrderPrintFromNode(t.root.Right)
	}
}

func (t *RBTree) preOrderPrintFromNode(node *Node) {
	if node.Left != nil {
		t.preOrderPrintFromNode(node.Left)
	}
	fmt.Println(node.Value)
	if node.Right != nil {
		t.preOrderPrintFromNode(node.Right)
	}
}

func (t *RBTree) insertToNode(node *Node, e interface{}) *Node {
	objectValue := node.Value.(int)
	if objectValue > e.(int) {
		if node.Left != nil {
			return t.insertToNode(node.Left, e)
		} else {
			newNode := Node{}
			newNode.Parent = node
			newNode.Value = e
			newNode.Color = "r"
			node.Left = &newNode
			return &newNode
		}
	} else {
		if node.Right != nil {
			return t.insertToNode(node.Right, e)
		} else {
			newNode := Node{}
			newNode.Parent = node
			newNode.Value = e
			newNode.Color = "r"
			node.Right = &newNode
			return &newNode
		}
	}

}

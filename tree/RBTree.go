package tree

type RBTree struct {
	root *Node
}

func (t *RBTree) insert(e interface{}) {
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
				t.insertToNode(t.root.Left, e)
			}
		} else {
			if t.root.Right == nil {
				newNode := Node{}
				newNode.Parent = t.root
				newNode.Value = e
				newNode.Color = "r"
				t.root.Right = &newNode
			} else {
				t.insertToNode(t.root.Right, e)
			}
		}

	}
}

func (t *RBTree) insertFixup(node *Node) {
	for node.Parent.Color == "r" {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Right
			if uncle.Color == "r" {
				node.Parent.Color = "b"
				uncle.Color = "b"
				node.Parent.Parent.Color = "r"
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {

				}
			}
		}
	}
}

func (t *RBTree) insertToNode(node *Node, e interface{}) {
	objectValue := node.Value.(int)
	if objectValue > e.(int) {
		if node.Left != nil {
			t.insertToNode(node.Left, e)
		} else {
			newNode := Node{}
			newNode.Parent = node
			newNode.Value = e
			newNode.Color = "r"
			node.Left = &newNode
		}
	} else {
		if node.Right != nil {
			t.insertToNode(node.Right, e)
		} else {
			newNode := Node{}
			newNode.Parent = node
			newNode.Value = e
			newNode.Color = "r"
			node.Right = &newNode
		}
	}

}

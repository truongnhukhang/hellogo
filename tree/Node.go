package tree

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  interface{}
}

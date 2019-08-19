package tree

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Color  string
	Value  interface{}
}

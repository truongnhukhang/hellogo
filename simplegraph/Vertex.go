package simplegraph

type Vertex struct {
	Color          string
	Value          string
	Index          int
	discoveredTime int
	finishedTime   int
	Distance       int
	Predecessors   *Vertex
}

func (v *Vertex) CompareWith(e interface{}) int {
	obj := e.(*Vertex)
	return v.Distance - obj.Distance
}

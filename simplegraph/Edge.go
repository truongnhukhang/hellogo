package simplegraph

import "strconv"

type Edge struct {
	Weight int
	value  int
	Source *Vertex
	Desc   *Vertex
}

func (e *Edge) CompareWith(edge interface{}) int {
	return e.Weight - edge.(*Edge).Weight
}

func (e *Edge) String() string {
	return e.Source.Value + "---" + e.Desc.Value + " : " + strconv.Itoa(e.Weight)
}

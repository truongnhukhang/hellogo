package simplegraph

import "strconv"

type Edge struct {
	Weight int
	value  int
	Source *Vertex
	Des    *Vertex
}

func (e *Edge) CompareWith(edge interface{}) int {
	return e.Weight - edge.(*Edge).Weight
}

func (e *Edge) String() string {
	return e.Source.Value + "---" + e.Des.Value + " : " + strconv.Itoa(e.Weight)
}

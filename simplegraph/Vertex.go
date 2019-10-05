package simplegraph

type Vertex struct {
	Color          string
	Value          string
	Distance       int
	Index          int
	Parent         *Vertex
	discoveredTime int
	finishedTime   int
}

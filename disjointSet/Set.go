package disjointSet

type Set struct {
	Key            interface{}
	Representative *Set
	Rank           int
}

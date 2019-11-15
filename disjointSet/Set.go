package disjointSet

type Set struct {
	key            interface{}
	representative *Set
	rank           int
}

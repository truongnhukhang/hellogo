package disjointSet

type DisjointSet struct {
	sets *map[interface{}]*Set
}

func (d *DisjointSet) init() *DisjointSet {
	sets := map[interface{}]*Set{}
	d.sets = &sets
	return d
}

func New() *DisjointSet {
	disjointSet := new(DisjointSet)
	return disjointSet.init()
}

func (d *DisjointSet) MakeSet(e interface{}) {
	set := Set{key: e}
	set.representative = &set
	set.rank = 0
	(*d.sets)[e] = &set
}

func (d *DisjointSet) findSet(e interface{}) {

}

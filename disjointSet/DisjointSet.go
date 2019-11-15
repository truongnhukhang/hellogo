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

func (d *DisjointSet) MakeSet(e interface{}) *Set {
	set := Set{Key: e}
	set.Representative = &set
	set.Rank = 0
	(*d.sets)[e] = &set
	return &set
}

func (d *DisjointSet) Union(e1 interface{}, e2 interface{}) {
	set1 := (*d.sets)[e1]
	set2 := (*d.sets)[e2]
	d.link(set1, set2)
}

func (d *DisjointSet) link(set1 *Set, set2 *Set) {
	set1 = d.findSet(set1)
	set2 = d.findSet(set2)
	if set1.Rank < set2.Rank {
		set1.Representative = set2
	} else {
		set2.Representative = set1
		if set2.Rank == set1.Rank {
			set2.Rank = set2.Rank + 1
		}
	}
}

func (d *DisjointSet) FindSet(e interface{}) *Set {
	set := (*d.sets)[e]
	return d.findSet(set)
}

func (d *DisjointSet) findSet(set *Set) *Set {
	if set != set.Representative {
		set.Representative = d.findSet(set.Representative)
	}
	return set.Representative
}

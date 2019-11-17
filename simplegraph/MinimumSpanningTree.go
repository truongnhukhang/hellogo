package simplegraph

import (
	"container/list"
	"github.com/truongnhukhang/hellogo/disjointSet"
	"github.com/truongnhukhang/hellogo/object"
	"github.com/truongnhukhang/hellogo/queue"
)

func FindMinimumSpanningTree(g *SimpleGraphL) *list.List {
	result := list.New()
	edgesDB := []object.Object{}
	priorityQueue := queue.PriorityQueue{edgesDB}
	disjoint := disjointSet.New()
	vertices := g.vertices
	for e := vertices.Front(); e != nil; e = e.Next() {
		disjoint.MakeSet(e.Value)
	}
	edges := g.edgesList
	for e := edges.Front(); e != nil; e = e.Next() {
		priorityQueue.Put(e.Value.(*Edge))
	}
	tempEdge := priorityQueue.Poll().(*Edge)
	for priorityQueue.Size() > 0 {
		if disjoint.FindSet(tempEdge.Source) != disjoint.FindSet(tempEdge.Des) {
			result.PushBack(tempEdge)
			disjoint.Union(tempEdge.Source, tempEdge.Des)
		}
		tempEdge = priorityQueue.Poll().(*Edge)
	}
	return result
}

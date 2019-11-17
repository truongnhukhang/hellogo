package simplegraph

import (
	"container/list"
	"github.com/truongnhukhang/hellogo/object"
	"github.com/truongnhukhang/hellogo/queue"
	"math"
)

func (g *SimpleGraphL) FindShortestPath(source string, destination string) *list.List {
	g.buildSingleSourceShortestPathDijkstra(source)
	destinationVertex := (*g.verticesValueMap)[destination]
	result := list.New()
	for destinationVertex.Predecessors != nil {
		result.PushFront(destinationVertex)
		destinationVertex = destinationVertex.Predecessors
	}
	return result
}

func (g *SimpleGraphL) buildSingleSourceShortestPathDijkstra(source string) {
	vertexSource := (*g.verticesValueMap)[source]
	initSingleSource(g, vertexSource)
	db := make([]object.Object, 0)
	vertexQueue := queue.PriorityQueue{DB: db}
	for e := g.vertices.Front(); e != nil; e = e.Next() {
		temp := e.Value.(*Vertex)
		vertexQueue.Put(temp)
	}
	var vertex *Vertex = nil
	var listVertices *list.List = nil
	for vertexQueue.Size() != 0 {
		vertex = vertexQueue.Poll().(*Vertex)
		listVertices = (*g.edgesAdj)[vertex]
		if listVertices != nil {
			for e := listVertices.Front(); e != nil; e = e.Next() {
				edge := e.Value.(*Edge)
				// relaxing
				if edge.Des.Distance > vertex.Distance+edge.Weight {
					edge.Des.Distance = vertex.Distance + edge.Weight
					edge.Des.Predecessors = vertex
					// ToDo : work around for decreaseKey of priorityQueue
					index := findIndexInArrayObj(vertexQueue.DB, edge.Des)
					if index != -1 {
						vertexQueue.DecreaseKey(index, edge.Des)
					}
				}
			}
		}
	}
}

func findIndexInArrayObj(objects []object.Object, vertex *Vertex) int {
	for i := 0; i < len(objects); i++ {
		if objects[i].(*Vertex).Value == vertex.Value {
			return i + 1
		}
	}
	return -1
}

func initSingleSource(g *SimpleGraphL, source *Vertex) {
	vertices := g.vertices
	for e := vertices.Front(); e != nil; e = e.Next() {
		temp := e.Value.(*Vertex)
		temp.Distance = math.MaxInt32
		temp.Predecessors = nil
	}
	source.Distance = 0
}

package main

import (
	"fmt"
	"github.com/truongnhukhang/hellogo/simplegraph"
	"strconv"
)

func dijkstra() {
	graph := simplegraph.New()
	a := graph.AddVertex("a")
	b := graph.AddVertex("b")
	c := graph.AddVertex("c")
	d := graph.AddVertex("d")
	e := graph.AddVertex("e")
	graph.AddEdgeAndWeight(a, b, 10)
	graph.AddEdgeAndWeight(a, c, 5)
	graph.AddEdgeAndWeight(b, c, 2)
	graph.AddEdgeAndWeight(b, e, 1)
	graph.AddEdgeAndWeight(c, b, 3)
	graph.AddEdgeAndWeight(c, e, 9)
	graph.AddEdgeAndWeight(c, d, 2)
	graph.AddEdgeAndWeight(d, a, 7)
	graph.AddEdgeAndWeight(d, e, 6)
	graph.AddEdgeAndWeight(e, d, 4)
	paths := graph.FindShortestPath("a", "e")
	pathString := "a"
	prv := 0
	for e := paths.Front(); e != nil; e = e.Next() {
		vertex := e.Value.(*simplegraph.Vertex)
		pathString = pathString + "--[" + strconv.Itoa(vertex.Distance-prv) + "]-->" + vertex.Value
		prv = vertex.Distance
	}
	fmt.Println(pathString)
}

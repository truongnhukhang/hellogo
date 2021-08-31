package main

import (
	"fmt"
	"github.com/truongnhukhang/hellogo/simplegraph"
)

func spanningTree() {
	graph := simplegraph.New()
	a := graph.AddVertex("a")
	b := graph.AddVertex("b")
	c := graph.AddVertex("c")
	d := graph.AddVertex("d")
	e := graph.AddVertex("e")
	f := graph.AddVertex("f")
	g := graph.AddVertex("g")
	h := graph.AddVertex("h")
	i := graph.AddVertex("i")

	graph.AddEdgeAndWeight(a, b, 4)
	graph.AddEdgeAndWeight(a, h, 8)
	graph.AddEdgeAndWeight(b, h, 11)
	graph.AddEdgeAndWeight(b, c, 8)
	graph.AddEdgeAndWeight(c, i, 2)
	graph.AddEdgeAndWeight(c, d, 7)
	graph.AddEdgeAndWeight(c, f, 4)
	graph.AddEdgeAndWeight(h, i, 7)
	graph.AddEdgeAndWeight(h, g, 1)
	graph.AddEdgeAndWeight(i, g, 6)
	graph.AddEdgeAndWeight(g, f, 2)
	graph.AddEdgeAndWeight(f, e, 10)
	graph.AddEdgeAndWeight(f, d, 13)
	graph.AddEdgeAndWeight(d, e, 9)

	result := simplegraph.FindMinimumSpanningTree(graph)
	for e := result.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*simplegraph.Edge).String())
	}

}

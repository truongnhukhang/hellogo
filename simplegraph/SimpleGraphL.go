package simplegraph

import (
	"container/list"
	"fmt"
)

type SimpleGraphL struct {
	vertices         *list.List
	verticesValueMap *map[string]*Vertex
	edges            *map[*Vertex]*list.List
}

func (g *SimpleGraphL) init() *SimpleGraphL {
	g.vertices = list.New()
	g.edges = &map[*Vertex]*list.List{}
	g.verticesValueMap = &map[string]*Vertex{}
	return g
}

func New() *SimpleGraphL {
	return new(SimpleGraphL).init()
}

func (g *SimpleGraphL) AddVertex(value string) *Vertex {
	vertex := Vertex{Value: value}
	inserted := g.vertices.PushBack(&vertex)
	(*g.verticesValueMap)[value] = inserted.Value.(*Vertex)
	return &vertex
}

func (g *SimpleGraphL) AddEdge(source *Vertex, desc *Vertex, weight int) {
	tempList := (*g.edges)[source]
	if tempList == nil {
		tempList = list.New()
		(*g.edges)[source] = tempList
	}
	tempList.PushBack(desc)
}

func (g *SimpleGraphL) PrintGraph() {
	fmt.Println("Vertices : ")
	node := g.vertices.Front()
	for node != nil {
		fmt.Print(" " + node.Value.(*Vertex).Value)
		node = node.Next()
	}
	fmt.Println(" ")
	fmt.Println("Edges : ")
	for key, value := range *g.edges {
		fmt.Print("[" + key.Value + "] ->")
		node = value.Front()
		for node != nil {
			fmt.Print(" " + node.Value.(*Vertex).Value)
			node = node.Next()
		}
		fmt.Println(" ")
	}
}

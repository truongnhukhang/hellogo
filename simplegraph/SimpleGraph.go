package simplegraph

import (
	"fmt"
	queue2 "github.com/truongnhukhang/hellogo/queue"
	"math"
	"strconv"
)

type SimpleGraph struct {
	NumVertex int
	vertices  []*Vertex
	edges     [][]*Edge
}

func (g *SimpleGraph) Init() {
	g.vertices = make([]*Vertex, 0, g.NumVertex)
	g.edges = make([][]*Edge, g.NumVertex)
	for i := 0; i < len(g.edges); i++ {
		g.edges[i] = make([]*Edge, g.NumVertex)
		for j := 0; j < len(g.edges[i]); j++ {
			g.edges[i][j] = &Edge{value: 0, Weight: 0}
		}
	}
}

func (g *SimpleGraph) AddVertex(value string) *Vertex {
	tempVertex := Vertex{Value: value}
	g.vertices = append(g.vertices, &tempVertex)
	tempVertex.Index = len(g.vertices) - 1
	return &tempVertex
}

func (g *SimpleGraph) AddEdge(vertexSource Vertex, vertexDes Vertex, weight int) {
	if vertexSource.Index < len(g.edges) && vertexDes.Index < len(g.edges) {
		g.edges[vertexSource.Index][vertexDes.Index].value = 1
		g.edges[vertexSource.Index][vertexDes.Index].Weight = weight
	}
}

func (g *SimpleGraph) PrintGraph() {
	for i := 0; i < len(g.vertices); i++ {
		fmt.Print(" " + g.vertices[i].Value + "-" + g.vertices[i].Color + "-" + strconv.Itoa(g.vertices[i].Distance))
	}
	fmt.Println("")
	fmt.Println("*******")
	fmt.Print(" ")
	for i := 0; i < len(g.vertices); i++ {
		fmt.Print(" " + g.vertices[i].Value)
	}
	fmt.Println("")
	for i := 0; i < len(g.vertices); i++ {
		fmt.Print(g.vertices[i].Value + " ")
		for j := 0; j < len(g.edges[i]); j++ {
			fmt.Print(strconv.Itoa(g.edges[i][j].value) + " ")
		}
		fmt.Println("")
	}

}

func (g *SimpleGraph) BreathFirstSearch(vertex *Vertex) {

	for i, _ := range g.vertices {
		g.vertices[i].Color = "white"
		g.vertices[i].Distance = math.MaxInt32
	}
	vertex.Color = "black"
	vertex.Distance = 0
	vertex.Parent = nil
	queue := queue2.SimpleQueue{}
	queue.Put(vertex)
	for !queue.IsEmpty() {
		e := queue.Poll().(*Vertex)
		for i, v := range g.edges[e.Index] {
			if v.value == 1 && g.vertices[i].Color == "white" {
				g.vertices[i].Parent = e
				g.vertices[i].Distance = e.Distance + 1
				queue.Put(g.vertices[i])
			}
		}
		e.Color = "black"
	}
}

package simplegraph

import (
	"container/list"
	"fmt"
	"strconv"
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
	vertex := Vertex{Value: value, Color: "white"}
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

func (g *SimpleGraphL) FindStrongConnectComponent() *list.List {
	reverseGraph := g.ReverseGraph()
	topologicalSortResult := g.TopologicalSort()
	// reset label data
	temp := topologicalSortResult.Front()
	for temp != nil {
		vertex := temp.Value.(*Vertex)
		vertex.Color = "white"
		temp = temp.Next()
	}
	// find root component
	result := list.New()
	temp = topologicalSortResult.Front()
	for temp != nil {
		vertex := temp.Value.(*Vertex)
		if vertex.Color == "white" {
			root := list.New()
			g.findStrongConnectComponentVisit(vertex, reverseGraph.edges, root)
			result.PushBack(root)
		}
		temp = temp.Next()
	}
	return result
}

func (g *SimpleGraphL) findStrongConnectComponentVisit(vertex *Vertex, edges *map[*Vertex]*list.List, root *list.List) {
	vertex.Color = "gray"
	root.PushBack(vertex)
	verticesDes := (*edges)[vertex]
	if verticesDes != nil {
		temp := verticesDes.Front()
		for temp != nil {
			tempVertex := temp.Value.(*Vertex)
			if tempVertex.Color == "white" {
				g.findStrongConnectComponentVisit(tempVertex, edges, root)
			}
			temp = temp.Next()
		}
	}
	vertex.Color = "black"
}

func (g *SimpleGraphL) ReverseGraph() *SimpleGraphL {
	reverseGraph := New()
	node := g.vertices.Front()
	for node != nil {
		reverseGraph.vertices.PushBack(node.Value)
		node = node.Next()
	}
	for key, value := range *g.edges {
		tempNode := value.Front()
		for tempNode != nil {
			reverseGraph.AddEdge(tempNode.Value.(*Vertex), key, 1)
			tempNode = tempNode.Next()
		}
	}
	return reverseGraph
}

func (g *SimpleGraphL) TopologicalSort() *list.List {
	sortResult := list.New()
	node := g.vertices.Front()
	time := 0
	for node != nil {
		vertex := node.Value.(*Vertex)
		if vertex.Color == "white" {
			g.topologicalSortVisit(vertex, sortResult, &time)
		}
		node = node.Next()
	}
	return sortResult
}

func (g *SimpleGraphL) topologicalSortVisit(vertex *Vertex, sortResult *list.List, time *int) {
	vertex.Color = "gray"
	*time = *time + 1
	vertex.discoveredTime = *time
	desVertices := (*g.edges)[vertex]
	if desVertices != nil {
		temp := desVertices.Front()
		for temp != nil {
			vertexDes := temp.Value.(*Vertex)
			if vertexDes.Color == "white" {
				g.topologicalSortVisit(vertexDes, sortResult, time)
			}
			temp = temp.Next()
		}
	}
	*time = *time + 1
	vertex.Color = "black"
	vertex.finishedTime = *time
	sortResult.PushFront(vertex)
}

func (g *SimpleGraphL) PrintGraph() {
	fmt.Println("Vertices : ")
	node := g.vertices.Front()
	for node != nil {
		vertex := node.Value.(*Vertex)
		fmt.Print(" " + vertex.Value + "(" + strconv.Itoa(vertex.discoveredTime) + "/" + strconv.Itoa(vertex.finishedTime) + ")")
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

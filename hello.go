package main

import (
	"container/list"
	"fmt"
	"github.com/truongnhukhang/hellogo/simplegraph"
)

func main() {
	//var a []int
	//
	//queue := MinPriorityQueue{a}
	//var data = []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1, 18}
	//for i := 0; i < len(data); i++ {
	//	queue.Put(data[i])
	//}
	//queue.Print()
	//fmt.Println("")
	//testPoll := queue.Poll()
	//for testPoll > math.MinInt32 {
	//	fmt.Println(testPoll)
	//	testPoll = queue.Poll()
	//}
	//queue.DecreaseKey(4, 4)
	//fmt.Println("")
	//queue.Print()
	//text := "khangggh"
	//hc := greedy.HuffmanCoding{}
	//hc.Context = text
	//compressData := hc.CompressData()
	//fmt.Println(compressData)
	//fmt.Println(hc.UncompressData(compressData))
	//testTree := tree.RBTree{}
	//testTree.Insert(4)
	//testTree.LevelOrderPrint()
	//testTree.Insert(8)
	//testTree.LevelOrderPrint()
	//testTree.Insert(5)
	//testTree.LevelOrderPrint()
	//testTree.Insert(12)
	//testTree.LevelOrderPrint()
	//testTree.Insert(6)
	//testTree.LevelOrderPrint()
	//testTree.Insert(10)
	//testTree.LevelOrderPrint()
	//testTree.Insert(14)
	//testTree.LevelOrderPrint()
	//testTree.Insert(9)
	//testTree.LevelOrderPrint()
	//testTree.Insert(11)
	//testTree.LevelOrderPrint()
	//testTree.Insert(13)
	//testTree.LevelOrderPrint()
	//testTree.Insert(15)
	//testTree.LevelOrderPrint()
	//graph := simplegraph.SimpleGraph{NumVertex: 4}
	//graph.Init()
	graph := simplegraph.New()
	a := graph.AddVertex("a")
	b := graph.AddVertex("b")
	c := graph.AddVertex("c")
	d := graph.AddVertex("d")
	e := graph.AddVertex("e")
	graph.AddEdge(a, b)
	graph.AddEdge(b, c)
	graph.AddEdge(c, a)
	graph.AddEdge(b, d)
	graph.AddEdge(d, e)
	graph.AddEdge(e, d)
	//sortResult := graph.TopologicalSort()
	//fmt.Println("sort result")
	//temp := sortResult.Front();
	//for temp!=nil {
	//	value := temp.Value.(*simplegraph.Vertex);
	//	fmt.Print(" "+value.Value);
	//	temp=temp.Next();
	//}
	strongComponent := graph.FindStrongConnectComponent()
	root := strongComponent.Front()
	for root != nil {
		temp := root.Value.(*list.List).Front()
		for temp != nil {
			value := temp.Value.(*simplegraph.Vertex)
			fmt.Print(" " + value.Value)
			temp = temp.Next()
		}
		fmt.Println("")
		root = root.Next()
	}
	graph.PrintGraph()
	graph.ReverseGraph().PrintGraph()
}

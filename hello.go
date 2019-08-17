package main

import (
	"github.com/truongnhukhang/hellogo/tree"
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
	//queue.IncreaseKey(4, 4)
	//fmt.Println("")
	//queue.Print()
	//text := "khangggh"
	//hc := greedy.HuffmanCoding{}
	//hc.Context = text
	//compressData := hc.CompressData()
	//fmt.Println(compressData)
	//fmt.Println(hc.UncompressData(compressData))
	bstTree := tree.BSTTree{}
	bstTree.Insert(8)
	bstTree.Insert(5)
	bstTree.Insert(12)
	bstTree.Insert(4)
	bstTree.Insert(6)
	bstTree.Insert(10)
	bstTree.Insert(14)
	bstTree.Insert(9)
	bstTree.Insert(11)
	bstTree.Insert(13)
	bstTree.Insert(15)
	bstTree.PreOrderPrint()
	node := bstTree.Search(12)
	bstTree.Delete(node)
	bstTree.PreOrderPrint()
}

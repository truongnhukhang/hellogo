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
	testTree := tree.RBTree{}
	testTree.Insert(4)
	testTree.LevelOrderPrint()
	testTree.Insert(8)
	testTree.LevelOrderPrint()
	testTree.Insert(5)
	testTree.LevelOrderPrint()
	testTree.Insert(12)
	testTree.LevelOrderPrint()
	testTree.Insert(6)
	testTree.LevelOrderPrint()
	testTree.Insert(10)
	testTree.LevelOrderPrint()
	testTree.Insert(14)
	testTree.LevelOrderPrint()
	testTree.Insert(9)
	testTree.LevelOrderPrint()
	testTree.Insert(11)
	testTree.LevelOrderPrint()
	testTree.Insert(13)
	testTree.LevelOrderPrint()
	testTree.Insert(15)
	testTree.LevelOrderPrint()
}

package main

import (
	"fmt"
	. "github.com/truongnhukhang/hellogo/queue"
	"math"
)

func main() {
	var a []int

	queue := MinPriorityQueue{a}
	var data = []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1, 18}
	for i := 0; i < len(data); i++ {
		queue.Put(data[i])
	}
	queue.Print()
	fmt.Println("")
	testPoll := queue.Poll()
	for testPoll > math.MinInt32 {
		fmt.Println(testPoll)
		testPoll = queue.Poll()
	}
	//queue.IncreaseKey(4, 4)
	//fmt.Println("")
	//queue.Print()

}

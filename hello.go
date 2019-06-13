package main

import (
	"fmt"
	. "github.com/truongnhukhang/hellogo/queue"
)

func main() {
	var a []int

	queue := MaxPriorityQueue{a}
	var data = []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1, 18}
	for i := 0; i < len(data); i++ {
		queue.Put(data[i])
	}
	queue.Print()
	queue.IncreaseKey(4, 4)
	fmt.Println("")
	queue.Print()

}

package queue

import (
	"fmt"
	. "github.com/truongnhukhang/hellogo/heap"
	"strconv"
)

type MinPriorityQueue struct {
	DB []int
}

func (q *MinPriorityQueue) Put(e int) {
	q.DB = append(q.DB, e)
	if len(q.DB) > 1 {
		ToMinHeap(q.DB)
	}
}

func (q *MinPriorityQueue) Poll() int {
	var e = q.DB[0]
	q.DB = append(q.DB[:0], q.DB[1:]...)
	if len(q.DB) > 1 {
		MinHeapBuild(q.DB, 1, len(q.DB))
	}
	return e
}

func (q *MinPriorityQueue) Min() int {
	var e = q.DB[0]
	return e
}

func (q *MinPriorityQueue) IncreaseKey(index int, value int) {
	q.DB[index-1] = value
	parent := index / 2
	for parent != 0 {
		if q.DB[parent-1] < q.DB[index-1] {
			q.DB[parent-1], q.DB[index-1] = q.DB[index-1], q.DB[parent-1]
			index = parent
		}
		parent = parent / 2
	}

}

func (q *MinPriorityQueue) Print() {
	for i := 0; i < len(q.DB); i++ {
		fmt.Print(strconv.Itoa(q.DB[i]) + " ")
	}
}

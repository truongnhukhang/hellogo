package queue

import (
	"fmt"
	. "github.com/truongnhukhang/hellogo/heap"
	"strconv"
)

type PriorityQueue struct {
	DB []int
}

func (q *PriorityQueue) Put(e int) {
	q.DB = append(q.DB, e)
	if len(q.DB) > 1 {
		ToMaxHeap(q.DB)
	}
}

func (q *PriorityQueue) Poll() interface{} {
	var e = q.DB[0]
	q.DB = append(q.DB[:0], q.DB[1:]...)
	if len(q.DB) > 1 {
		MaxHeapBuild(q.DB, 1, len(q.DB))
	}
	return e
}

func (q *PriorityQueue) Max() interface{} {
	var e = q.DB[0]
	return e
}

func (q *PriorityQueue) IncreaseKey(index int, value int) {
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

func (q *PriorityQueue) Print() {
	for i := 0; i < len(q.DB); i++ {
		fmt.Print(strconv.Itoa(q.DB[i]) + " ")
	}
}

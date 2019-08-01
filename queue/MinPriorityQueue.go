package queue

import (
	"fmt"
	. "github.com/truongnhukhang/hellogo/heap"
	"math"
	"strconv"
)

type MinPriorityQueue struct {
	DB []int
}

func (q *MinPriorityQueue) Put(e int) {
	q.DB = append([]int{e}, q.DB...)
	if len(q.DB) > 1 {
		MinHeapBuild(q.DB, 1, len(q.DB))
	}
}

func (q *MinPriorityQueue) Poll() int {
	if len(q.DB) > 0 {
		var e = q.DB[0]
		q.DB = append(q.DB[:0], q.DB[1:]...)
		if len(q.DB) > 1 {
			MinHeapBuild(q.DB, 1, len(q.DB))
		}
		return e
	}
	return math.MinInt32
}

func (q *MinPriorityQueue) Peek() int {
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

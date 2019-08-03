package queue

import (
	"fmt"
	"github.com/truongnhukhang/hellogo/heap"
	"github.com/truongnhukhang/hellogo/object"
)

type PriorityQueue struct {
	DB []object.Object
}

func (q *PriorityQueue) Put(e object.Object) {
	q.DB = append([]object.Object{e}, q.DB...)
	if len(q.DB) > 1 {
		heap.MinHeapModify(q.DB, 1, len(q.DB))
	}
}

func (q *PriorityQueue) Poll() object.Object {
	if len(q.DB) > 0 {
		var e = q.DB[0]
		q.DB = append(q.DB[:0], q.DB[1:]...)
		if len(q.DB) > 1 {
			heap.MinHeapModify(q.DB, 1, len(q.DB))
		}
		return e
	}
	return nil
}

func (q *PriorityQueue) Peek() object.Object {
	var e = q.DB[0]
	return e
}

func (q *PriorityQueue) Print() {
	for i := 0; i < len(q.DB); i++ {
		fmt.Println(q.DB[i])
	}
}

func (q *PriorityQueue) Size() int {
	return len(q.DB)
}

package queue

import (
	"fmt"
	"github.com/truongnhukhang/hellogo/object"
)

type PriorityQueue struct {
	DB []object.Object
}

func (q *PriorityQueue) Put(e object.Object) {
	q.DB = append([]object.Object{e}, q.DB...)
	if len(q.DB) > 1 {
		q.minHeapModify(1, len(q.DB))
	}
}

func (q *PriorityQueue) Poll() object.Object {
	if len(q.DB) > 0 {
		var e = q.DB[0]
		q.DB = append(q.DB[:0], q.DB[1:]...)
		if len(q.DB) > 1 {
			q.minHeapModify(1, len(q.DB))
		}
		return e
	}
	return nil
}

func (q *PriorityQueue) DecreaseKey(index int, value object.Object) {
	q.DB[index-1] = value
	parent := index / 2
	for parent != 0 {
		if q.DB[parent-1].CompareWith(q.DB[index-1]) > 0 {
			q.DB[parent-1], q.DB[index-1] = q.DB[index-1], q.DB[parent-1]
			index = parent
		}
		parent = parent / 2
	}
}

func (q *PriorityQueue) minHeapModify(index int, length int) {
	left := 2 * index
	right := 2*index + 1
	minimum := index
	if left <= length && q.DB[left-1].CompareWith(q.DB[minimum-1]) < 0 {
		minimum = left
	}
	if right <= length && q.DB[right-1].CompareWith(q.DB[minimum-1]) < 0 {
		minimum = right
	}
	if minimum != index {
		q.DB[minimum-1], q.DB[index-1] = q.DB[index-1], q.DB[minimum-1]
		q.minHeapModify(minimum, length)
	}
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

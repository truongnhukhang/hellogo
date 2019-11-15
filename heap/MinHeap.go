package heap

import "github.com/truongnhukhang/hellogo/object"

func MinHeapModify(a []object.Object, index int, length int) {
	var left = 2 * index
	var right = 2*index + 1
	var minium = index
	if left <= length && a[index-1].CompareWith(a[left-1]) > 0 {
		minium = left
	}
	if right <= length && a[minium-1].CompareWith(a[right-1]) > 0 {
		minium = right
	}
	if minium != index {
		a[index-1], a[minium-1] = a[minium-1], a[index-1]
		MinHeapModify(a, minium, length)
	}
}

package heap

import "github.com/truongnhukhang/hellogo/object"

func MinHeapModify(a []object.Object, index int, length int) {
	var left = 2 * index
	var right = 2*index + 1
	var minimum = index
	if left <= length && a[index-1].CompareWith(a[left-1]) > 0 {
		minimum = left
	}
	if right <= length && a[minimum-1].CompareWith(a[right-1]) > 0 {
		minimum = right
	}
	if minimum != index {
		a[index-1], a[minimum-1] = a[minimum-1], a[index-1]
		MinHeapModify(a, minimum, length)
	}
}

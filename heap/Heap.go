package heap

func ToMaxHeap(a []int) {
	for i := len(a) / 2; i > 0; i-- {
		MaxHeapBuild(a, i, len(a))
	}
}
func MaxHeapBuild(a []int, index int, length int) {
	var left = 2 * index
	var right = 2*index + 1
	var largest = index
	if left <= length && a[index-1] < a[left-1] {
		largest = left
	}
	if right <= length && a[largest-1] < a[right-1] {
		largest = right
	}
	if largest != index {
		swapValue(a, index-1, largest-1)
		MaxHeapBuild(a, largest, length)
	}
}

func swapValue(a []int, index1 int, index2 int) {
	var temp = a[index1]
	a[index1] = a[index2]
	a[index2] = temp
}

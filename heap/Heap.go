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
func ToMinHeap(a []int) {
	for i := len(a) / 2; i > 0; i-- {
		MinHeapBuild(a, i, len(a))
	}
}
func MinHeapBuild(a []int, index int, length int) {
	var left = 2 * index
	var right = 2*index + 1
	var minium = index
	if left <= length && a[index-1] > a[left-1] {
		minium = left
	}
	if right <= length && a[minium-1] > a[right-1] {
		minium = right
	}
	if minium != index {
		swapValue(a, index-1, minium-1)
		MinHeapBuild(a, minium, length)
	}
}

func InsertToMinHeap(a []int, value int, index int) {
	a[index-1] = value
	parent := index / 2
	for parent != 0 {
		if a[index-1] < a[parent-1] {
			a[parent-1], a[index-1] = a[index-1], a[parent-1]
			index = parent
		}
		parent = parent / 2
	}

}

func swapValue(a []int, index1 int, index2 int) {
	var temp = a[index1]
	a[index1] = a[index2]
	a[index2] = temp
}

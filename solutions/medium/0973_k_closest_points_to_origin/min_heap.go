package kclosest

type Heap struct {
	arr [][]int
}

func (h *Heap) Pop() []int {
	if len(h.arr) == 0 {
		return nil
	}

	x := h.arr[0]

	lastIndex := len(h.arr) - 1

	h.arr[0] = h.arr[lastIndex]
	h.arr = h.arr[:lastIndex]

	if len(h.arr) > 0 {
		h.Heapify(0)
	}

	return x
}

func BuildHeap(input [][]int) *Heap {
	heap := &Heap{arr: input}

	for idx := len(input)/2 - 1; idx > -1; idx-- {
		heap.Heapify(idx)
	}

	return heap
}

func (h *Heap) Heapify(idx int) {
	smallest := idx
	lc := idx*2 + 1
	rc := idx*2 + 2

	if lc < len(h.arr) && distance(h.arr[lc]) < distance(h.arr[smallest]) {
		smallest = lc
	}

	if rc < len(h.arr) && distance(h.arr[rc]) < distance(h.arr[smallest]) {
		smallest = rc
	}

	if smallest != idx {
		h.arr[idx], h.arr[smallest] = h.arr[smallest], h.arr[idx]
		h.Heapify(smallest)
	}
}

func distance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

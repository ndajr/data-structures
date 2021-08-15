package datastructures

type Heap struct {
	items []int
}

func (h *Heap) Init(items []int) {
	// TODO: use heapify for better performance
	for _, v := range items {
		h.add(v)
	}
}

func (h Heap) Len() int {
	return len(h.items)
}

func (h Heap) leftIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h Heap) rightIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h Heap) parentIndex(childIndex int) int {
	if childIndex-1 < 0 {
		return -1
	}
	return (childIndex - 1) / 2
}

func (h Heap) left(idx int) int {
	l := h.leftIndex(idx)
	if l > h.Len()-1 {
		return -1
	}
	return h.items[l]
}

func (h Heap) right(idx int) int {
	r := h.rightIndex(idx)
	if r > h.Len()-1 {
		return -1
	}
	return h.items[r]
}

func (h Heap) parent(idx int) int {
	if h.parentIndex(idx) == -1 {
		return -1
	}
	return h.items[h.parentIndex(idx)]
}

func (h Heap) hasLeft(idx int) bool {
	if h.left(idx) == -1 {
		return false
	}
	return h.leftIndex(idx) < h.Len()
}

func (h Heap) hasRight(idx int) bool {
	if h.right(idx) == -1 {
		return false
	}
	return h.rightIndex(idx) < h.Len()
}

func (h Heap) hasParent(idx int) bool {
	p := h.parent(idx)
	if p == -1 {
		return false
	}
	return p >= 0
}

func (h *Heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h Heap) Peek() int {
	if h.Len() == 0 {
		return -1
	}
	return h.items[0]
}

func (h *Heap) Poll() bool {
	if h.Len() == 0 {
		return false
	}
	last := h.Len() - 1
	h.items[0] = h.items[last]
	h.items = h.items[:last]
	h.down()
	return true
}

func (h *Heap) add(item int) {
	h.items = append(h.items, item)
	h.up()
}

func (h *Heap) up() {
	idx := h.Len() - 1
	for h.hasParent(idx) && h.parent(idx) > h.items[idx] {
		h.swap(h.parentIndex(idx), idx)
		idx = h.parentIndex(idx)
	}
}

func (h *Heap) down() {
	idx := 0
	for h.hasLeft(idx) {
		smallerChildIndex := h.leftIndex(idx)
		if h.hasRight(idx) && h.right(idx) < h.left(idx) {
			smallerChildIndex = h.rightIndex(idx)
		}

		if h.items[idx] < h.items[smallerChildIndex] {
			break
		} else {
			h.swap(idx, smallerChildIndex)
		}

		idx = smallerChildIndex
	}
}

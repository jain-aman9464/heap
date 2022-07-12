package main

import "fmt"

type minHeap struct {
	heapArr  []int
	size     int
	capacity int
}

func newMinHeap(maxsize int) *minHeap {
	return &minHeap{
		heapArr:  make([]int, 0, maxsize),
		size:     0,
		capacity: maxsize,
	}
}

func (m *minHeap) parent(i int) int {
	return (i - 1) / 2
}

func (m *minHeap) left(i int) int {
	return 2*i + 1
}

func (m *minHeap) right(i int) int {
	return 2*i + 2
}

func (m *minHeap) swap(first, second int) {
	temp := m.heapArr[first]
	m.heapArr[first] = m.heapArr[second]
	m.heapArr[second] = temp
}

func (m *minHeap) insert(value int) error {
	if m.size >= m.capacity {
		return fmt.Errorf("Heap is full, can't insert")
	}

	m.heapArr = append(m.heapArr, value)
	m.size++
	idx := m.size - 1
	for m.heapArr[idx] < m.heapArr[m.parent(idx)] {
		m.swap(idx, m.parent(idx))
	}

	return nil
}

func (m *minHeap) heapify(idx int) {
	smallest := idx
	leftChild := m.left(idx)
	rightChild := m.right(idx)

	if leftChild < m.size && m.heapArr[leftChild] < m.heapArr[smallest] {
		smallest = leftChild
	}

	if rightChild < m.size && m.heapArr[rightChild] < m.heapArr[smallest] {
		smallest = rightChild
	}

	if smallest != idx {
		m.swap(smallest, idx)
		m.heapify(smallest)
	}
}

func (m *minHeap) buildMinHeap() {
	for index := ((m.size / 2) - 1); index >= 0; index-- {
		m.heapify(index)
	}
}

func (m *minHeap) extractMin() int {
	if m.size <= 0 {
		return -1
	}

	root := m.heapArr[0]
	m.heapArr[0] = m.heapArr[m.size-1]
	m.size--
	m.heapify(0)
	return root
}

func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	minHeap := newMinHeap(len(inputArray))
	for i := 0; i < len(inputArray); i++ {
		minHeap.insert(inputArray[i])
	}
	minHeap.buildMinHeap()

	fmt.Println("min: ", minHeap.extractMin())
	fmt.Println("min: ", minHeap.extractMin())
	fmt.Println("min: ", minHeap.extractMin())
	fmt.Println("min: ", minHeap.extractMin())
	fmt.Println("min: ", minHeap.extractMin())
	fmt.Println("min: ", minHeap.extractMin())
}

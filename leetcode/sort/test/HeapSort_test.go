package test

import (
	"testing"
)

func TestInitHeap(t *testing.T) {
	heap := InitHeap()
	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(7)
	heap.removeMax()

}

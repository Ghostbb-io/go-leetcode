package _703

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	tmp := old[len(*h)-1]
	// 切片重切，因為 heap 套件在內部實作時是靠
	// 切片的長度在定位的，跟我們先前的方式不太一樣
	// 所以一定要重切
	*h = old[0 : len(*h)-1]
	return tmp
}

type KthLargest struct {
	k  int
	hp *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	intHeap := IntHeap(nums)
	kt := KthLargest{k, &intHeap}
	heap.Init(kt.hp)
	return kt
}

func (k *KthLargest) Add(val int) int {
	heap.Push(k.hp, val)
	for k.hp.Len() > k.k {
		heap.Pop(k.hp)
	}
	return (*k.hp)[0]
}

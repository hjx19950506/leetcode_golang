package mergeKLists

import "container/heap"
/*
type ListNode struct {
	Val int
	Next *ListNode
}
*/
//使用官方包"container/heap"
//为ListNodePtrHeap实现heap接口的五个函数
//在heap中保存每个链表的当前节点
//每当pop出heap的最小根节点时
//将该节点所在链表的下一个节点push入heap
type ListNodePtrHeap []*ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	lnph := ListNodePtrHeap(make([]*ListNode, 0))
	for _, v := range lists {
		if v != nil {
			lnph = append(lnph, v)
		}
	}
	h := &lnph
	heap.Init(h)

	headAhead := ListNode{-1, nil}
	res, ptr := &headAhead, &headAhead
	for h.Len() > 0 {
		ptr.Next = heap.Pop(h).(*ListNode)
		ptr = ptr.Next
		if ptr.Next != nil {
			heap.Push(h, ptr.Next)
		}
	}

	return res.Next
}

func (h ListNodePtrHeap) Len() int {
	return len(h)
}

func (h ListNodePtrHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodePtrHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodePtrHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodePtrHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
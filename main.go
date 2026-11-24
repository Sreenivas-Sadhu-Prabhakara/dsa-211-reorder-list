package main

import "fmt"

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// Reorder rearranges L0→L1→…→Ln into L0→Ln→L1→Ln-1→… in place: find the
// middle, reverse the back half, then interleave the two halves.
func Reorder(head *Node) {
	if head == nil || head.Next == nil {
		return
	}
	// Middle (end of first half): slow stops before the second half.
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	second := slow.Next
	slow.Next = nil // split the halves

	// Reverse the second half.
	var prev *Node
	for cur := second; cur != nil; {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	// Interleave: first half is never shorter than the reversed second.
	first := head
	for prev != nil {
		fNext, sNext := first.Next, prev.Next
		first.Next = prev
		prev.Next = fNext
		first = fNext
		prev = sNext
	}
}

func fromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(head *Node) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func main() {
	l := fromSlice([]int{1, 2, 3, 4, 5})
	Reorder(l)
	fmt.Println(toSlice(l))
}

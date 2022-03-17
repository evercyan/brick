package helper

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// Slice2Node ...
func Slice2Node(nums []int) *ListNode {
	head := &ListNode{}
	tmp := head
	for _, v := range nums {
		tmp.Next = &ListNode{Val: v}
		tmp = tmp.Next
	}
	return head.Next
}

// Node2Slice ...
func Node2Slice(l *ListNode) []int {
	nums := make([]int, 0)
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	return nums
}

// Slice2Cycle ...
func Slice2Cycle(nums []int, pos int) *ListNode {
	head := Slice2Node(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}

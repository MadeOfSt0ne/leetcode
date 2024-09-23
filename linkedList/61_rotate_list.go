package linkedlist

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lenght := 1
	curr := head
	for curr.Next != nil {
		lenght += 1
		curr = curr.Next
	}
	curr.Next = head

	k %= lenght
	tail := lenght - k - 1
	curr = head

	for i := 0; i < tail; i++ {
		curr = curr.Next
	}
	result := curr.Next
	curr.Next = nil
	return result
}

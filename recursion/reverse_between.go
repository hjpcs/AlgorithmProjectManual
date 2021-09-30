package recursion

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// base case
	if left == 1 {
		// 相当于反转前 right 个元素
		reverseN(head, right)
	}
	// 前进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var successor *ListNode //后驱结点

// 反转以 head 为起点的 n 个结点，返回新的头结点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 记录第 n+1 个结点
		successor = head.Next
		return head
	}
	// 以 head.Next 为起点，需要反转前 n-1 个结点
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	// 让反转后的 head 结点和后面的结点连起来
	head.Next = successor
	return last
}

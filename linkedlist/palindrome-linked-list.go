package linkedlist

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/11/11 2:41 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func IsPalindrome(head *ListNode) bool {
	var slow *ListNode = head
	var fast *ListNode = head
	var temp *ListNode = nil
	var prev *ListNode = nil
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		temp = slow.next
		slow.next = prev
		prev = slow
		slow = temp
	}
	if fast != nil || fast.next != nil {
		slow = slow.next
	}
	var l1 *ListNode = prev
	var l2 *ListNode = slow
	for l1 != nil && l2 != nil && l1.value == l2.value {
		l1 = l1.next
		l2 = l2.next
	}
	return l1 == nil && l2 == nil
}

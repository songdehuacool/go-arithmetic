package _7_linkedlist

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/12 10:05 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func NewListNode(value interface{}, node *ListNode) *ListNode {
	return &ListNode{
		Value: value,
		Next:  node,
	}
}

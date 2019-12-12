package _7_linkedlist

import "fmt"

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/12 10:08 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
type LinkedList struct {
	Head   *ListNode
	Length int32
}

func NewLinkedList(value interface{}) *LinkedList {
	return &LinkedList{NewListNode(value, nil), 0}
}

// 查找指定节点
func (this *LinkedList) Find(value interface{}) *ListNode {
	if this.Head == nil {
		return nil
	}

	p := this.Head
	for nil != p {
		if p.Value == value {
			return p
		}
		p = p.Next
	}
	return nil
}

// 尾部插入
func (this *LinkedList) InsertToTail(value interface{}) {
	// 判断节点是否为空
	if this == nil {
		this = NewLinkedList(value)
	}
	// 首先判断头节点是否为空
	if this.Head == nil {
		// 创建新节点
		head := NewListNode(value, nil)
		this.Head = head
		return
	}
	// for循环找到尾节点
	// 判断条件 nil != p.Next 保证最后p节点不会为空
	p := this.Head
	for nil != p.Next {
		p = p.Next
	}
	newNode := NewListNode(value, nil)
	p.Next = newNode
}

// 指定节点后面插入
func (this *LinkedList) InsertToNodeAfter(value interface{}, node *ListNode) {
	p := this.Head
	for nil != p {
		if p == node {
			break
		}
		p = p.Next
	}
	// 没有此节点
	if p == nil {
		return
	}
	newNode := NewListNode(value, nil)
	newNode.Next = node.Next
	node.Next = newNode
}

// 指定节点前面插入
func (this *LinkedList) InsertToNodeBefore(value interface{}, node *ListNode) {
	if this.Head == nil {
		return
	}

	p := this.Head
	// 找到待插入节点的前驱节点
	for nil != p.Next {
		if p.Next.Value == node.Value {
			break
		}
		p = p.Next
	}
	if p == nil {
		return
	}
	newNode := NewListNode(value, nil)
	newNode.Next = p.Next
	p.Next = newNode
}

// 头节点插入
func (this *LinkedList) InsertToHead(value interface{}) {
	if this.Head == nil {
		return
	}
	newNode := NewListNode(value, nil)
	newNode.Next = this.Head.Next
	this.Head.Next = newNode
}

func (this *LinkedList) DeleteByNode(node *ListNode) {
	if this.Head == node {
		return
	}

	p := this.Head
	// 找到前驱节点
	for nil != p.Next {
		// 找到前驱节点
		if p.Next.Value == node.Value {
			break
		}
		p = p.Next
	}
	if p == nil {
		return
	}
	// 开始执行删除操作
	p.Next = p.Next.Next
}

func (this *LinkedList) Print() {
	p := this.Head
	for nil != p {
		fmt.Printf("%v--->", p.Value)
		p = p.Next
	}
}

func InversionLinkedList_head(node *ListNode) {
	list := NewLinkedList(9999)
	head := list.Head
	// 头节点与下一节点连接
	head.Next = node
	cur := node.Next
	var next *ListNode = nil
	for nil != cur {
		next = cur.Next
		cur.Next = head.Next
		head.Next = cur
		cur = next
	}
}

// 1-->2-->3-->4-->5-->6
// 1.保存当前指针的下一指针
// 2.将当前指针指向前驱指针
// 3.当前节点赋值给前驱节点
// 4.下一指针赋值给当前节点
func InversionLinkedList(node *ListNode) *ListNode {
	// 当前节点
	cur := node
	// 前驱节点
	var pre *ListNode = nil
	for nil != cur {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp

	}
	return pre
}

func Compare(node1, node2 *ListNode) bool {
	for node1 != nil && node2 != nil {
		if node1.Value != node2.Value {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	if nil != node1 || nil != node2 {
		return false
	}
	return true
}

func Plalindrome(node *ListNode) bool {
	if node == nil {
		return false
	}
	if node.Next == nil {
		return true
	}

	slow := node
	fast := node
	for nil != slow.Next && nil != fast.Next.Next {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var rightNode *ListNode = nil
	var leftNode *ListNode = nil
	if fast.Next == nil {
		// 奇数个节点 slow即为中间节点
		rightNode = slow.Next
		leftNode = InversionLinkedListPlalindrome(node, slow).Next
	} else {
		// 偶数个节点
		rightNode = slow.Next
		leftNode = InversionLinkedListPlalindrome(node, slow)
	}
	return Compare(leftNode, rightNode)
}

func InversionLinkedListPlalindrome(head *ListNode, node *ListNode) *ListNode {
	// 当前节点
	cur := head
	// 前驱节点
	var pre *ListNode = nil
	for cur != node {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp

	}
	// 单独处理node节点
	node.Next = pre
	return node
}

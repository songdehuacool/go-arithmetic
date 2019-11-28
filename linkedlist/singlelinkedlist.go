package linkedlist

import "fmt"

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/11/10 3:18 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length int
}

func NewListNode(value interface{}) *ListNode {
	return &ListNode{
		next:  nil,
		value: value,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   NewListNode(0),
		length: 0,
	}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

// 在p节点后面插入数据
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	// 1.检查需要插入的数据是否为空
	if p == nil {
		return false
	}
	// 2.新建一个节点
	newNode := NewListNode(v)
	// 3.将p节点指向的节点保存起来
	oldNext := p.next
	// 4.将p的指针指向新节点
	p.next = newNode
	// 5.将新节点指向旧的节点
	newNode.next = oldNext
	// 6.链表长度加1
	this.length++
	return true
}

// 在p节点前面插入新节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	// 1.判断数据是否合法，不能等于nil或者头节点
	if p == nil || p == this.head {
		return false
	}
	// 1.循环找到p点,从头节点都下一个节点开始进行循环查找
	cur := this.head.next
	// 前驱节点
	pre := this.head
	// 从第一个开始遍历，直到最后cur为nil。如果中间找到直接break掉for循环
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	// for循环执行到最后没有找到p点
	if nil == cur {
		return false
	}

	// 2.创建一个新节点
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

// 在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if int(index) >= this.length {
		return nil
	}
	// 头指针不作为开始节点
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除传入的节点
// 1.判断需要操作的节点是否合法
// 2.找到要操作的节点在链表中的位置
// 3.删除节点
// 4链表长度减一
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	pre.next = cur.next
	p = nil
	this.length--
	return true
}

// 打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

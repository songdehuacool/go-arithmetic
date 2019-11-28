package linkedlist

import (
	"testing"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/11/11 2:59 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func TestIsPalindrome(t *testing.T) {
	strs := []string{"123", "", "", "", "123"}
	l := NewLinkedList()
	for _, v := range strs {
		l.InsertToTail(v)
	}
	l.Print()
	t.Log(IsPalindrome(l.head.next))
}

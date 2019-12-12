package _7_linkedlist

import (
	"fmt"
	"testing"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/12 2:05 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

func TestLinkedList_(t *testing.T) {
	list := NewLinkedList(9999)
	list.InsertToTail(1)
	list.InsertToTail(3)
	list.Print()
	fmt.Println()
	list.InsertToHead(0)
	list.Print()
	fmt.Println()
	list.InsertToNodeBefore(2, list.Head.Next.Next.Next)
	list.Print()
	fmt.Println()
}

func TestLinkedList_InversionLinkedList(t *testing.T) {
	list := NewLinkedList(0)
	list.InsertToTail(1)
	list.InsertToTail(2)
	list.InsertToTail(2)
	list.InsertToTail(1)
	list.InsertToTail(0)
	//list.Print()
	//fmt.Println()
	//node := InversionLinkedList(list.Head)
	//fmt.Println("node ------------ ", node.Value, "--->", node.Next.Value, "--->", node.Next.Next.Value)

	fmt.Println("------------------", Plalindrome(list.Head))
}

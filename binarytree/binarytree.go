package binarytree

import "fmt"

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/11/28 9:39 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
// 树节点
type TreeNode struct {
	value       interface{}
	left, right *TreeNode
}

// 先序遍历
func (this *TreeNode) PreOrderTraversal() {
	if this == nil {
		return
	}
	fmt.Printf("%v-->", this.value)
	this.left.PreOrderTraversal()
	this.right.PreOrderTraversal()
}

// 中序遍历
func (this *TreeNode) InOrderTraversal() {
	if this == nil {
		return
	}
	this.left.InOrderTraversal()
	fmt.Printf("%v-->", this.value)
	this.right.InOrderTraversal()
}

// 后序遍历
func (this *TreeNode) PostOrderTraversal() {
	if this == nil {
		return
	}
	this.left.PostOrderTraversal()
	fmt.Printf("%v-->", this.value)
	this.right.PostOrderTraversal()
}

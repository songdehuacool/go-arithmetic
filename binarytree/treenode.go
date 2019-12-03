package binarytree

import "fmt"

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/3 6:34 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
type TreeNode struct {
	data        interface{}
	left, right *TreeNode
}

func NewTreeNode(v interface{}) *TreeNode {
	return &TreeNode{data: v}
}

func (this *TreeNode) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", this.data, this.left, this.right)
}

// 前序遍历
func (node *TreeNode) preOrderTraverse() {
	if node == nil {
		return
	}
	fmt.Println("-->", node.data)
	node.left.preOrderTraverse()
	node.right.preOrderTraverse()
}

// 中序遍历
func (node *TreeNode) inOrderTraverse() {
	if node == nil {
		return
	}
	node.left.inOrderTraverse()
	fmt.Println("-->", node.data)
	node.right.inOrderTraverse()
}

// 后序遍历
func (node *TreeNode) postOrderTraverse() {
	if node == nil {
		return
	}
	node.left.postOrderTraverse()
	node.right.postOrderTraverse()
	fmt.Println("-->", node.data)
}

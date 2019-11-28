package binarytree

import (
	"fmt"
	"testing"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/11/28 10:35 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

func TestTreeNode_Traversal(t *testing.T) {
	//创建一颗树
	root := TreeNode{"A", nil, nil}
	root.left = &TreeNode{value: "B"}
	root.right = &TreeNode{value: "C"}
	root.left.left = &TreeNode{value: "D"}
	root.left.right = &TreeNode{value: "F"}
	root.left.right.left = new(TreeNode)
	root.left.right.left.value = "E"

	root.right.left = &TreeNode{value: "G"}
	root.right.left.right = &TreeNode{value: "H"}
	root.right.right = &TreeNode{value: "I"}
	root.PreOrderTraversal()
	fmt.Println("\n****************************************")
	root.InOrderTraversal()
	fmt.Println("\n****************************************")
	root.PostOrderTraversal()
}

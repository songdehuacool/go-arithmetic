package binarytree

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/5 9:56 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
type BSF struct {
	*BinaryTree
	compareFun func(v, vNode interface{}) int
}

func NewBSF(value int, compareFun func(v, vNode interface{}) int) *BSF {
	if compareFun == nil {
		return nil
	}
	return &BSF{NewBinaryTree(value), compareFun}
}

func (this *BSF) Find(v interface{}) *TreeNode {
	p := this.root
	for nil != p {
		compResult := this.compareFun(v, p.data)
		if compResult > 0 {
			p = p.right
		} else if compResult < 0 {
			p = p.left
		} else {
			return p
		}
	}
	return nil
}

// 新增节点
func (this *BSF) Insert(v interface{}) bool {
	p := this.root
	for nil != p {
		compResult := this.compareFun(v, p.data)
		if compResult == 0 {
			return false
		} else if compResult > 0 {
			if nil == p.right {
				p.right = NewTreeNode(v)
				break
			}
			p = p.right
		} else {
			if nil == p.left {
				p.left = NewTreeNode(v)
				break
			}
			p = p.left
		}
	}
	return true
}

func (this *BSF) Delete(data interface{}) {
	p := this.root
	var pp *TreeNode = nil
	for nil != p && p.data != data {
		pp = p
		if data.(int) > p.data.(int) {
			p = p.right
		} else {
			p = p.left
		}
	}

	if p == nil {
		return
	}

	if p.left != nil && p.right != nil {
		minP := p.right
		minPP := p
		for minP.left != nil {
			minPP = minP
			minP = minP.left
		}
		// 将minP中的值替换到p中
		p.data = minP.data
		// 将minP的引用指向p
		p = minP
		pp = minPP
	}

	// 删除节点是叶子节点或者仅有一个子节点
	var child *TreeNode
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}

	if pp == nil {
		p = child
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}
}

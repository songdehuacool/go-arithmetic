package binarytree

import (
	"testing"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/5 11:19 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

var compareFunc = func(v, nodeV interface{}) int {
	vInt := v.(int)
	nodeVInt := nodeV.(int)
	if vInt > nodeVInt {
		return 1
	} else if vInt < nodeVInt {
		return -1
	}
	return 0
}

func TestBSF_Find(t *testing.T) {
	bsf := NewBSF(0, compareFunc)
	bsf.Insert(1)
	bsf.Insert(4)
	bsf.Insert(2)
	bsf.Insert(3)
	bsf.Insert(5)
	bsf.Delete(3)
	t.Log(bsf.Find(3))
}

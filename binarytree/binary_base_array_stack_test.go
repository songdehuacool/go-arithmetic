package binarytree

import (
	"testing"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/4 10:21 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

func TestNewArrayStack(t *testing.T) {
	arryaStack := NewArrayStack(0)
	arryaStack.push(1)
	arryaStack.push(2)
	arryaStack.push(3)
	arryaStack.Print()
	arryaStack.pop()
	arryaStack.Print()

}

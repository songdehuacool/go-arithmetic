package binarytree

import "fmt"

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/4 10:00 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
type ArrayStack struct {
	top  int
	data []interface{}
}

func NewArrayStack(value interface{}) *ArrayStack {
	return &ArrayStack{-1, make([]interface{}, 0, 32)}
}

func (this *ArrayStack) isEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

func (this *ArrayStack) push(value interface{}) {
	if this.isEmpty() {
		this.top = 0
	} else {
		this.top += 1
	}
	if this.top > len(this.data)-1 {
		this.data = append(this.data, value)
	} else {
		this.data[this.top] = value
	}
}

func (this *ArrayStack) pop() interface{} {
	if this.isEmpty() {
		return nil
	}
	v := this.data[this.top]
	this.top -= 1
	return v
}

func (this *ArrayStack) getTop() interface{} {
	if this.isEmpty() {
		return nil
	}
	return this.data[this.top]
}

func (this *ArrayStack) flush() {
	this.top = -1
}

func (this *ArrayStack) Print() {
	if this.isEmpty() {
		fmt.Println("stack is empty")
	} else {
		for i := this.top; i >= 0; i-- {
			fmt.Println(this.data[i])
		}
	}
}

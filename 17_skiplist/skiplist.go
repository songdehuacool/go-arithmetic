package _7_skiplist

import (
	"math"
	"math/rand"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/11/26 11:57 上午
 * @description：
 * @modified By：
 * @version    ：$
 */


const (
	MAX_LEVEL = 16
)

// 跳表节点
type SkipListNode struct {
	// 当前高度
	h int32
	// 分值
	score int32
	// 存储值
	value interface{}
	// 当前节点保存的不同级数的下一跳节点
	forwars []*SkipListNode
}

// 跳表
type SkipList struct {
	// 头节点
	head *SkipListNode
	// 长度
	length int32
	// 高度
	level  int32
}

func NewSkipListNode(value interface{}, s int32, level int32) *SkipListNode {
	return &SkipListNode{
		h:       level,
		score:   s,
		value:   value,
		forwars: make([]*SkipListNode, level, level),
	}
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:   NewSkipListNode(0, math.MaxInt32, MAX_LEVEL),
		length: 0,
		level:  1,
	}
}

func (skip *SkipList) Insert(v interface{}, score int32) int32 {
	if skip == nil || v == nil {
		return -1
	}
	i := MAX_LEVEL - 1
	cur := skip.head
	update := [MAX_LEVEL]*SkipListNode{}
	// 找到需要插入位置的节点
	for ; i >= 0; i-- {
		for nil != cur.forwars[i] {
			if cur.value != nil {
				if cur.forwars[i].value == v {
					return -2
				}
				if cur.forwars[i].score > score {
					update[i] = cur
					break
				}
				cur = cur.forwars[i]
			}
			if cur.forwars[i] == nil {
				update[i] = cur
			}
		}
	}
	var h int32 = 1
	// 随机生成层高
	for ; h <= MAX_LEVEL; h++ {
		if rand.Int31() % 7 == 1 {
			h++
		}
	}

	// 生成新节点
	newSkipListNode := NewSkipListNode(v, score, h)

	// 插入到原有节点
	for i := 0; i <= len(update); i++ {
		oleNode := update[i].forwars[i]
		update[i].forwars[i] = newSkipListNode
		newSkipListNode.forwars[i] = oleNode
	}
	if h > skip.level {
		skip.level = h
	}
	skip.length++
	return 0
}

func (this *SkipList) Find(v interface{}, score int32) *SkipListNode {
	if this == nil || v == nil {
		return nil
	}
	i := MAX_LEVEL - 1
	cur := this.head
	for ; i >= 0; i-- {
		for cur.forwars[i] != nil {
			if cur.forwars[i].value == v && cur.forwars[i].score == score {
				return cur
			} else if cur.forwars[i].score > score {
				break
			}
			cur = cur.forwars[i]
		}
	}
	return nil
}





















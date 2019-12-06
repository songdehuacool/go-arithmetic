package heap

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/6 12:04 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

type Heap struct {
	data   []int32
	length int32
	count  int32
}

func NewHeap(capacity int32) *Heap {
	return &Heap{make([]int32, 0, capacity+1), capacity, 0}
}

func (this *Heap) Insert(value int32) {
	if this.count >= this.length {
		return
	}
	this.count++
	this.data[this.count] = value
	for i := this.count; i/2 > 0 && this.data[i] > this.data[i/2]; {
		this.swap(i, i/2)
		i = i / 2
	}
}

func (this *Heap) RemoveMax() {
	if this.count < 1 {
		return
	}
	this.data[1] = this.data[this.count]
	this.count--
	this.heapify(1)
}

// 自顶向下堆化 因为每个节点有可能有两个子树所以需要验证2 * i和 2 * i + 1
// 与自下向上堆化区别在于，每个子节点只有一个父亲节点，所以只需要 i / 2即可
func (this *Heap) heapify(i int32) {
	for {
		maxPos := i
		// 2 * i < this.count 从数组起始位置开始按2的倍数向后递增,小于等于当前存储个数
		if 2*i <= this.count && this.data[i] < this.data[i*2] {
			maxPos = i * 2
		}

		if 2*i+1 <= this.count && this.data[maxPos] < this.data[i*2+1] {
			maxPos = i*2 + 1
		}

		if i == maxPos {
			break
		}
		this.swap(i, maxPos)
		i = maxPos
	}
}

func (this *Heap) Build(arr []int32, length int32) {
	this.data = arr
	for i := length / 2; i >= 1; i-- {
		this.heapify(i)
	}
}

func (this *Heap) Sort(arr []int32, n int32) {
	this.Build(arr, n)
	k := n
	for k > 1 {
		this.swap(1, k)
		k--
		this.heapify(k)
	}
}

func (this *Heap) swap(v1, v2 int32) {
	temp := this.data[v1]
	this.data[v1] = this.data[v2]
	this.data[v2] = temp
}

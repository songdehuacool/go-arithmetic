package sort

import "testing"

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/9/12 3:37 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func TestMergeSort(t *testing.T) {
	arr := []int{5, 4}
	MergeSort(arr)
	t.Log(arr)

	arr = []int{5, 4, 3, 2, 1}
	MergeSort(arr)
	t.Log(arr)
}

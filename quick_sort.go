package sort

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/9/16 4:01 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func QuickSort(arr []int) {
	arrLen := len(arr)
	quickSort(arr, 0, arrLen-1)
}

func quickSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	// i已经排序好是中间节点，不再参与排序
	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

func partition(arr []int, start int, end int) int {
	// 选择最后一个数据作为比较值
	pivot := arr[end]
	// i标记已排序的数据
	i := start
	// 最后一个数据为比较值，不用比较所以此处为
	// j < end
	for j := start; j < end; j++ {
		// 如果arr[j]小于比较值，则将
		// arr[i]与arr[j]交换位置
		if arr[j] < pivot {
			if !(i == j) {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	// 经过以上排序后，i指向的是当前数组中数据     // 的最大值
	// 将最后一个值与arr[i]交换
	arr[i], arr[end] = arr[end], arr[i]
	return i
}

package sort

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/9/12 12:04 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	mergeSort(arr, 0, arrLen-1)
}

//递归调用
func mergeSort(arr []int, start int, end int) {
	//递归结束条件
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	//将数组合并
	merge(arr, start, mid, end)
}

//将两个有序数组合并成一个有序数组
func merge(arr []int, start int, mid int, end int) {
	tmpArr := make([]int, (end-start)+1)
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		//如果i指向的数据小于j指向的数据
		//则将i指向的数据复制到临时数组中
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}
	//查找两个分组中剩余数据
	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}
	//将临时数组值拷贝回原数组
	copy(arr[start:end+1], tmpArr)
}

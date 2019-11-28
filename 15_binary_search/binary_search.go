package _5_binary_search

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/11/16 8:17 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

func bsearch(arr []int, value int) int {
	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == value {
			return mid
		} else if arr[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

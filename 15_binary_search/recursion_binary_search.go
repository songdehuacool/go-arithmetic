package _5_binary_search

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/11/16 10:06 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func recursionBinarySearch(arr []int, value int) int  {
	return recursion(arr, value, 0, len(arr) - 1)
}

func recursion(arr []int, value, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + ((high - low) >> 1)
	if arr[mid] == value {
		return mid
	} else if arr[mid] > value {
		high = mid - 1
		return recursion(arr, value, low, high)
	} else {
		low = mid + 1
		return recursion(arr, value, low, high)
	}
}
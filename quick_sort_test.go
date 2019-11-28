package sort

import (
	"log"
	"testing"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/9/16 3:59 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func TestQuickSort(t *testing.T) {
	arr := []int{6, 2, 10, 1, 9}
	QuickSort(arr)
	log.Println(arr)

	arr2 := []int{6, 6}
	QuickSort(arr2)
	log.Println(arr2)
}

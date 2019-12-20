package leetcode

import (
	"fmt"
	"testing"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/17 10:46 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

func TestLeetCode_18(t *testing.T) {
	// [1,0,-1,0,-2,2]
	// nums := []int{1, 0, -1, 0, -2, 2}
	nums := []int{0, 0, 0, 0}
	b := fourSum(nums, 0)
	fmt.Println(b)
}

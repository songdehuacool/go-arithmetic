package leetcode

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/9/23 11:08 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
import "math"

func maxArea(height []int) int {
	var maxArea float64
	l := 0
	r := len(height) - 1
	for l < r {
		maxArea = math.Max(float64(maxArea), math.Min(float64(height[l]), float64(height[r]))*float64((r-l)))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return int(maxArea)
}

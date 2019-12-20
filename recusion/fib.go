package recusion

import (
	"fmt"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/18 4:58 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

func Fib(n int) int {
	// 不带缓存
	//num := recursion(n)
	//fmt.Println(num)
	//return num

	// 带缓存
	// 缓存数组
	memo := make([]int, n+1)
	num := recursion(n, memo)
	fmt.Println(num)
	return num
}

//func recursion(n int) int {
//	if n <= 1 {
//		return n
//	}
//	return recursion(n - 1) + recursion(n - 2)
//}

func recursion(n int, memo []int) int {
	if n <= 1 {
		return n
	}
	if memo[n] == 0 {
		memo[n] = recursion(n-1, memo) + recursion(n-2, memo)
		fmt.Println(memo)
		fmt.Println(memo[n])
		return memo[n]
	}
	return memo[n]
}

func Test1234567890() {

}

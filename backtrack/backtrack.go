package backtrack

import (
	"fmt"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/12/19 6:55 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	// 生成一个n x n的二维数组，并赋值"."
	board := make([]string, n)
	for i := 0; i < n; i++ {
		item := make([]byte, n)
		for j := 0; j < n; j++ {
			item[j] = '.'
		}
		board[i] = string(item[:])
	}
	backtrack(board, &res, n, 0)
	return res
}

func backtrack(board []string, res *[][]string, n, row int) {
	// 结束条件
	if row == n {
		*res = append(*res, board)
		return
	}
	// 列循环
	for col := 0; col < n; col++ {
		// 去除不合法的数据
		if !isOk(board, n, row, col) {
			continue
		}
		chars := []byte(board[row])
		chars[col] = 'Q'
		board[row] = string(chars)
		// 选择分支,行数➕1
		backtrack(board, res, n, row+1)
		// 恢复到之前的数据
		chars = []byte(board[row])
		chars[col] = '.'
		board[row] = string(chars)
	}
}

func isOk(board []string, n int, row int, col int) bool {
	fmt.Println(board)
	// 查询上方数据
	for i := 0; i < row; i++ {
		char := fmt.Sprintf("%c", []byte(board[i])[col])
		if char == "Q" {
			return false
		}
	}
	// 查询左上角
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		char := fmt.Sprintf("%c", []byte(board[row])[col])
		if char == "Q" {
			return false
		}
	}
	// 查询右上角
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		char := fmt.Sprintf("%c", []byte(board[row])[col])
		if char == "Q" {
			return false
		}
	}
	return true
}

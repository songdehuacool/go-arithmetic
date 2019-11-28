package _5_binary_search

import "testing"

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2019/11/16 9:57 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

func Test_bsearch(t *testing.T) {
	type args struct {
		arr   []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_bsearch",
			args{
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				5,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bsearch(tt.args.arr, tt.args.value); got != tt.want {
				t.Errorf("bsearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
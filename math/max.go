package math

import "github.com/zhouyusd/algorithm"

// Max 返回a和b的较大者
func Max[T algorithm.Comparable](a, b T) (min T) {
	if a > b {
		return a
	}
	return b
}

// Max1 返回a和b的较大者
// cmp 自定义类型比较器，若 a < b 应返回 -1，若 a = b 应返回 0，若 a > b 应返回 1
func Max1[T any](a, b T, cmp algorithm.Comparator[T]) (min T) {
	if cmp(a, b) > 0 {
		return a
	}
	return b
}

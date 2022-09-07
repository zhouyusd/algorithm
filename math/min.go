package math

import "github.com/zhouyusd/algorithm"

// Min 返回a和b的较小者
func Min[T algorithm.Comparable](a, b T) (min T) {
	if a > b {
		return b
	}
	return a
}

// Min1 返回a和b的较小者
// cmp 自定义类型比较器，若 a < b 应返回 -1，若 a = b 应返回 0，若 a > b 应返回 1
func Min1[T any](a, b T, cmp algorithm.Comparator[T]) (min T) {
	if cmp(a, b) > 0 {
		return b
	}
	return a
}

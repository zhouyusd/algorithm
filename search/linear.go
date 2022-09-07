package search

import "github.com/zhouyusd/algorithm"

// Linear 简单线性搜索算法，在最坏情况下迭代数组的所有元素
// 如果 x 在 a 中出现，则返回 x 在 a 中的下标，否则返回 -1 和 ErrNotFound
func Linear[T algorithm.Comparable](a []T, x T) (int, error) {
	for i, item := range a {
		if item == x {
			return i, nil
		}
	}
	return -1, ErrNotFound
}

// LinearX 忽略 Linear error
func LinearX[T algorithm.Comparable](a []T, x T) int {
	i, _ := Linear(a, x)
	return i
}

// Linear1 如果 x 在 a 中出现，则返回 x 在 a 中的下标，否则返回 -1 和 ErrNotFound
// cmp 自定义类型比较器，若 a < b 应返回 -1，若 a = b 应返回 0，若 a > b 应返回 1
func Linear1[T any](a []T, x T, cmp algorithm.Comparator[T]) (int, error) {
	for i, item := range a {
		if cmp(item, x) == 0 {
			return i, nil
		}
	}
	return -1, ErrNotFound
}

// Linear1X 忽略 Linear1 error
func Linear1X[T any](a []T, x T, cmp algorithm.Comparator[T]) int {
	i, _ := Linear1(a, x, cmp)
	return i
}

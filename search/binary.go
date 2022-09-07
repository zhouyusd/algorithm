package search

import "github.com/zhouyusd/algorithm"

// Binary 采用二分查找 如果 x 在 a 中出现，则返回 x 在 a 中的下标，否则返回 -1 和 ErrNotFound
func Binary[T algorithm.Comparable](a []T, x T) (int, error) {
	s, e := 0, len(a)-1
	for s <= e {
		mid := s + (e-s)/2
		if a[mid] > x {
			e = mid - 1
		} else if a[mid] < x {
			s = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, ErrNotFound
}

// BinaryX 忽略 Binary error
func BinaryX[T algorithm.Comparable](a []T, x T) int {
	i, _ := Binary(a, x)
	return i
}

// Binary1 采用二分查找 如果 x 在 a 中出现，则返回 x 在 a 中的下标，否则返回 -1 和 ErrNotFound
// cmp 自定义类型比较器，若 a < b 应返回 -1，若 a = b 应返回 0，若 a > b 应返回 1
func Binary1[T any](a []T, x T, cmp algorithm.Comparator[T]) (int, error) {
	s, e := 0, len(a)-1
	for s <= e {
		mid := s + (e-s)/2
		t := cmp(a[mid], x)
		if t > 0 {
			e = mid - 1
		} else if t < 0 {
			s = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, ErrNotFound
}

// Binary1X 忽略 Binary1 error
func Binary1X[T any](a []T, x T, cmp algorithm.Comparator[T]) int {
	i, _ := Binary1(a, x, cmp)
	return i
}

// LowerBound 返回范围[0，len（数组）-1]中不小于（即大于或等于）目标的第一个元素的索引
// 如果未找到此类元素，则返回 -1 和 ErrNotFound
func LowerBound[T algorithm.Comparable](a []T, x T) (int, error) {
	s, e := 0, len(a)-1
	for s <= e {
		mid := s + (e-s)/2
		if a[mid] < x {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	if s >= len(a) {
		return -1, ErrNotFound
	}
	return s, nil
}

// LowerBoundX 忽略 LowerBound error
func LowerBoundX[T algorithm.Comparable](a []T, x T) int {
	i, _ := LowerBound(a, x)
	return i
}

// LowerBound1 返回范围[0，len（数组）-1]中不小于（即大于或等于）目标的第一个元素的索引
// 如果未找到此类元素，则返回 -1 和 ErrNotFound
// cmp 自定义类型比较器，若 a < b 应返回 -1，若 a = b 应返回 0，若 a > b 应返回 1
func LowerBound1[T any](a []T, x T, cmp algorithm.Comparator[T]) (int, error) {
	s, e := 0, len(a)-1
	for s <= e {
		mid := s + (e-s)/2
		if cmp(a[mid], x) < 0 {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	if s >= len(a) {
		return -1, ErrNotFound
	}
	return s, nil
}

// LowerBound1X 忽略 LowerBound1 error
func LowerBound1X[T any](a []T, x T, cmp algorithm.Comparator[T]) int {
	i, _ := LowerBound1(a, x, cmp)
	return i
}

// UpperBound 返回范围[lowIndex，len（数组）-1]中大于目标的第一个元素的索引
// 如果未找到此类元素，则返回 -1 和 ErrNotFound
func UpperBound[T algorithm.Comparable](a []T, x T) (int, error) {
	s, e := 0, len(a)-1
	for s <= e {
		mid := s + (e-s)/2
		if a[mid] > x {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	if s >= len(a) {
		return -1, ErrNotFound
	}
	return s, nil
}

// UpperBoundX 忽略 UpperBound error
func UpperBoundX[T algorithm.Comparable](a []T, x T) int {
	i, _ := UpperBound(a, x)
	return i
}

// UpperBound1 返回范围[lowIndex，len（数组）-1]中大于目标的第一个元素的索引
// 如果未找到此类元素，则返回 -1 和 ErrNotFound
// cmp 自定义类型比较器，若 a < b 应返回 -1，若 a = b 应返回 0，若 a > b 应返回 1
func UpperBound1[T any](a []T, x T, cmp algorithm.Comparator[T]) (int, error) {
	s, e := 0, len(a)-1
	for s <= e {
		mid := s + (e-s)/2
		if cmp(a[mid], x) > 0 {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	if s >= len(a) {
		return -1, ErrNotFound
	}
	return s, nil
}

// UpperBound1X 忽略 UpperBound1 error
func UpperBound1X[T any](a []T, x T, cmp algorithm.Comparator[T]) int {
	i, _ := UpperBound1(a, x, cmp)
	return i
}

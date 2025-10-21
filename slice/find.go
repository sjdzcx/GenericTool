package slice

// Find 寻找一个,如果存在则返回下标
func Find[T comparable](s []T, target T) (int, bool) {
	for i, v := range s {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

// FindFunc 寻找一个满足条件的,如果存在则返回下标
func FindFunc[T any](s []T, f func(T) bool) (int, bool) {
	for i, v := range s {
		if f(v) {
			return i, true
		}
	}
	return -1, false
}

// FindFuncAll 寻找所有满足条件的,如果存在则返回切片
func FindFuncAll[T any](s []T, f func(T) bool) ([]T, bool) {
	result := make([]T, 5)
	for _, v := range s {
		if f(v) {
			result = append(result, v)
			//return i, true
		}
	}
	return nil, false
}

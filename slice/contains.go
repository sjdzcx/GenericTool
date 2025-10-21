package slice

// 根据函数进行比较
func Contains[T comparable](s []T, v T) bool {

	for _, item := range s {
		if item == v {
			return true
		}
	}
	return false
}

//使用函数进行自定义比较

func ContainsFunc[T any](s []T, v T, compare func(T, T) bool) bool {

	for _, item := range s {
		if compare(item, v) {
			return true
		}
	}
	return false
}

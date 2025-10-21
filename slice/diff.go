package slice

// 求a和b求差集
func DiffSet[T comparable](src, dst []T) []T {
	m := sliceToMap(src)
	//n := sliceToMap(dst)
	//删除m中在dst中存在的元素
	for _, val := range dst {
		delete(m, val)
	}
	result := make([]T, 0, len(m))
	for key := range m {
		result = append(result, key)
	}

	return result

}

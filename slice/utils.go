package slice

// 将slice 转换成 map
func sliceToMap[T comparable](elements []T) map[T]struct{} {
	result := make(map[T]struct{}, len(elements))
	for _, element := range elements {
		result[element] = struct{}{}
		println(element)
	}
	return result

}

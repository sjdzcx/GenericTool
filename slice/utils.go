package slice

// 将slice 转换成 map
func sliceToMap[T comparable](elements []T) map[T]struct{} {
	result := make(map[T]struct{}, len(elements))
	for _, element := range elements {
		result[element] = struct{}{}
	}
	return result

}
func deduplicate[T comparable](data []T) []T {
	m := sliceToMap(data)
	result := make([]T, 0, len(m))
	for key := range m {
		result = append(result, key)
	}
	return result
}

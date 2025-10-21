package slice

func IntersectSet[T comparable](a, b []T) []T {
	m := sliceToMap(a)
	result := make([]T, 0, 8)
	for _, v := range b {
		if _, found := m[v]; found {
			result = append(result, v)
		}
	}
	return deduplicate(result)
}

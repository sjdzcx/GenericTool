package slice

func Union[T comparable](slices ...[]T) []T {
	result := []T{}
	seen := make(map[T]struct{})

	for _, slice := range slices {
		for _, item := range slice {
			if _, exists := seen[item]; !exists {
				seen[item] = struct{}{}
				result = append(result, item)
			}
		}
	}
	return result
}

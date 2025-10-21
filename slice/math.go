package slice

import ct "github.com/sjdzcx/GenericTool/internal/constraint"

func Max[T ct.Number](s []T) (max T, ok bool) {
	if len(s) == 0 {
		return max, false
	}
	max = s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max, true
}
func Min[T ct.Number](s []T) (min T, ok bool) {
	if len(s) == 0 {
		return min, false
	}
	min = s[0]
	for _, v := range s[1:] {
		if v < min {
			min = v
		}
	}
	return min, true
}

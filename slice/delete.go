package slice

import "github.com/sjdzcx/GenericTool/internal/slice"

//删除Index中位置的元素

func Delete[T any](s []T, index int) ([]T, error) {
	res, _, err := slice.Delete(s, index)
	return res, err
}

// 在原切片上进行切改
func DeleteFilter[T any](s []T, filter func(int, T) bool) ([]T, error) {
	index := 0
	for idx, val := range s {
		if !filter(idx, val) {
			continue
		}
		s[index] = val
		index++
	}
	return s[:index], nil
}

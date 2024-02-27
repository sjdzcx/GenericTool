package slice

import (
	"github.com/sjdzcx/GenericTool/internal/err"
)

// Delete 删除指定下标的元素
func Delete[T any](src []T, index int) ([]T, T, error) {
	var ret T
	length := len(src)
	if index >= length || index < 0 {
		return src, ret, err.NewErrIndexOutOfRange(length, index)
	}
	ret = src[index]
	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	src = src[:length-1]
	return src, ret, nil

}

func DeleteElement[T comparable](src []T, element T) ([]T, error) {
	for i := 0; i < len(src); i++ {
		if src[i] == element {
			t, _, e := Delete(src, i)
			return t, e
		}

	}
	return src, nil

}

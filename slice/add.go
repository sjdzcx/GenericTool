package slice

import (
	"GenericTool/err"
)

// Add 在index 位置出添加元素
func Add[T any](src []T, element T, index int) ([]T, error) {
	if index > len(src) || index < 0 {
		return src, err.NewErrIndexOutOfRange(len(src), index)
	}
	r1 := make([]T, 0, len(src)+1)
	for i := 0; i < index; i++ {
		r1 = append(r1, src[i])
	}
	r1 = append(r1, element)
	for i := index; i < len(src); i++ {
		r1 = append(r1, src[i])
	}
	src = r1
	//src = append(r1, src[index:]...)
	return src, nil
}

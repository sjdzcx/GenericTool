package slice

import (
	"github.com/sjdzcx/GenericTool/internal/slice"
)

func Add[T any](src []T, element T, index int) ([]T, error) {
	return slice.Add(src, element, index)
}

package slice

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	tests := []struct {
		name   string
		inputs [][]int
		want   []int
	}{
		{"空切片", [][]int{}, []int{}},
		{"单个切片", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"无交集", [][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{"部分交集", [][]int{{1, 2}, {2, 3}}, []int{1, 2, 3}},
		{"完全相同", [][]int{{1, 2}, {1, 2}}, []int{1, 2}},
		{"重复元素", [][]int{{1, 2, 2}, {2, 3, 3}}, []int{1, 2, 3}},
		{"多个切片", [][]int{{1}, {2}, {3}, {1, 2}}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Union(tt.inputs...)
			gotMap := make(map[int]struct{})
			for _, v := range got {
				gotMap[v] = struct{}{}
			}
			wantMap := make(map[int]struct{})
			for _, v := range tt.want {
				wantMap[v] = struct{}{}
			}
			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

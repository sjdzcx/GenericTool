package slice

import (
	"reflect"
	"testing"
)

func TestDiffSet(t *testing.T) {

	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{"空切片", []int{}, []int{}, []int{}},
		{"无交集", []int{1, 2, 3}, []int{4, 5}, []int{1, 2, 3}},
		{"部分交集", []int{1, 2, 3}, []int{2, 4}, []int{1, 3}},
		{"完全相同", []int{1, 2}, []int{1, 2}, []int{}},
		{"重复元素", []int{1, 2, 2, 3}, []int{2}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DiffSet(tt.src, tt.dst)
			// 因为结果无序，转为map比较
			gotMap := make(map[int]struct{})
			for _, v := range got {
				gotMap[v] = struct{}{}
			}
			wantMap := make(map[int]struct{})
			for _, v := range tt.want {
				wantMap[v] = struct{}{}
			}
			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("diffSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

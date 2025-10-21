package slice

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	tests := []struct {
		name    string
		src     []int
		index   int
		want    []int
		wantErr bool
	}{
		{"正常删除中间", []int{1, 2, 3}, 1, []int{1, 3}, false},
		{"删除头部", []int{1, 2, 3}, 0, []int{2, 3}, false},
		{"删除尾部", []int{1, 2, 3}, 2, []int{1, 2}, false},
		{"越界删除", []int{1, 2, 3}, 3, []int{1, 2, 3}, true},
		{"负索引", []int{1, 2, 3}, -1, []int{1, 2, 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Delete(tt.src, tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFilter(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	tests := []struct {
		name   string
		filter func(int, int) bool
		want   []int
	}{
		{"删除偶数", func(_ int, v int) bool { return v%2 != 0 }, []int{1, 3, 5}},
		{"全部保留", func(_ int, v int) bool { return true }, []int{1, 2, 3, 4, 5}},
		{"全部删除", func(_ int, v int) bool { return false }, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteFilter(append([]int{}, src...), tt.filter)
			if err != nil {
				t.Errorf("DeleteFilter() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

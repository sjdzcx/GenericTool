package slice

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{"空切片", []int{}, []int{}},
		{"单元素", []int{1}, []int{1}},
		{"偶数长度", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"奇数长度", []int{1, 2, 3}, []int{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseInPlace(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{"空切片", []int{}, []int{}},
		{"单元素", []int{1}, []int{1}},
		{"偶数长度", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"奇数长度", []int{1, 2, 3}, []int{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := make([]int, len(tt.s))
			copy(s, tt.s)
			ReverseInPlace(s)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("ReverseInPlace() = %v, want %v", s, tt.want)
			}
		})
	}
}

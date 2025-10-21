package slice

import (
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want int
		ok   bool
	}{
		{"空切片", []int{}, 0, false},
		{"正常", []int{1, 3, 2}, 3, true},
		{"重复最大值", []int{5, 5, 2}, 5, true},
		{"负数", []int{-1, -3, -2}, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Max(tt.src)
			if ok != tt.ok || got != tt.want {
				t.Errorf("Max() = %v, %v; want %v, %v", got, ok, tt.want, tt.ok)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want int
		ok   bool
	}{
		{"空切片", []int{}, 0, false},
		{"正常", []int{1, 3, 2}, 1, true},
		{"重复最小值", []int{2, 2, 5}, 2, true},
		{"负数", []int{-1, -3, -2}, -3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Min(tt.src)
			if ok != tt.ok || got != tt.want {
				t.Errorf("Min() = %v, %v; want %v, %v", got, ok, tt.want, tt.ok)
			}
		})
	}
}

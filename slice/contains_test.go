package slice

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		v    int
		want bool
	}{
		{"空切片", []int{}, 1, false},
		{"存在元素", []int{1, 2, 3}, 2, true},
		{"不存在元素", []int{1, 2, 3}, 4, false},
		{"重复元素", []int{1, 2, 2, 3}, 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.s, tt.v)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsFunc(t *testing.T) {
	compare := func(a, b string) bool { return len(a) == len(b) }
	tests := []struct {
		name    string
		s       []string
		v       string
		compare func(string, string) bool
		want    bool
	}{
		{"空切片", []string{}, "a", compare, false},
		{"长度相等", []string{"ab", "cd", "efg"}, "xy", compare, true},
		{"长度不等", []string{"abc", "de"}, "fghij", compare, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsFunc(tt.s, tt.v, tt.compare)
			if got != tt.want {
				t.Errorf("ContainsFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

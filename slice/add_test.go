package slice

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		src     []int
		element int
		index   int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "中间插入",
			args:    args{src: []int{1, 2, 3}, element: 9, index: 1},
			want:    []int{1, 9, 2, 3},
			wantErr: false,
		},
		{
			name:    "头部插入",
			args:    args{src: []int{1, 2, 3}, element: 8, index: 0},
			want:    []int{8, 1, 2, 3},
			wantErr: false,
		},
		{
			name:    "尾部插入",
			args:    args{src: []int{1, 2, 3}, element: 7, index: 3},
			want:    []int{1, 2, 3, 7},
			wantErr: false,
		},
		{
			name:    "越界插入",
			args:    args{src: []int{1, 2, 3}, element: 6, index: 4},
			want:    []int{1, 2, 3},
			wantErr: true,
		},
		{
			name:    "负数索引",
			args:    args{src: []int{1, 2, 3}, element: 5, index: -1},
			want:    []int{1, 2, 3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.src, tt.args.element, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

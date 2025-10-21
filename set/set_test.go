package set

import (
	"reflect"
	"testing"
)

func TestMapSet_Add_Exists_Delete_Size_Values(t *testing.T) {
	s := NewMapSet[int](0)

	// 测试 Add 和 Exists
	s.Add(1)
	s.Add(2)
	if !s.Exists(1) || !s.Exists(2) {
		t.Errorf("Add 或 Exists 方法有误")
	}
	if s.Exists(3) {
		t.Errorf("Exists 方法有误，错误地返回了存在")
	}

	// 测试 Size
	if s.Size() != 2 {
		t.Errorf("Size 方法有误，期望 2，实际 %d", s.Size())
	}

	// 测试 Values
	values := s.Values()
	expected := []int{1, 2}
	// 因为 map 无序，这里用 reflect.DeepEqual 检查内容
	if !reflect.DeepEqual(map[int]struct{}{values[0]: {}, values[1]: {}}, map[int]struct{}{expected[0]: {}, expected[1]: {}}) {
		t.Errorf("Values 方法有误，期望 %v，实际 %v", expected, values)
	}

	// 测试 Delete
	s.Delete(1)
	if s.Exists(1) {
		t.Errorf("Delete 方法有误，元素未被删除")
	}
	if s.Size() != 1 {
		t.Errorf("Delete 后 Size 方法有误，期望 1，实际 %d", s.Size())
	}
}

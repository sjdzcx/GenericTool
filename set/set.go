package set

type Set[T comparable] interface {
	Add(value T)
	Delete(value T)
	Exists(value T) bool
	Size() int
	Values() []T
}

type MapSet[T comparable] struct {
	m map[T]struct{}
}

func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}
func (m *MapSet[T]) Add(value T) {
	//	panic("implement me")
	m.m[value] = struct{}{}
}

func (m *MapSet[T]) Delete(value T) {
	delete(m.m, value)
}

func (m *MapSet[T]) Exists(value T) bool {
	_, exists := m.m[value]
	return exists
}

func (m *MapSet[T]) Size() int {
	return len(m.m)
}

func (m *MapSet[T]) Values() []T {
	values := make([]T, 0, len(m.m))
	for key := range m.m {
		values = append(values, key)
	}
	return values
}

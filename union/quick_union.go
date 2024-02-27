package union

type QuickUnion struct {
	arr  []int
	size []int
}

func (q *QuickUnion) GetSize() int {
	cnt := 0
	for i := range q.arr {
		if q.arr[i] == i {
			cnt++
		}
	}
	return cnt

}

func NewQuickUnion(n int) *QuickUnion {
	arr := make([]int, n+1)
	size := make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i] = i
		size[i] = 1
	}

	return &QuickUnion{arr: arr, size: size}
}
func (q *QuickUnion) Connection(a, b int) bool {
	return q.Find(a) == q.Find(b)
}

func (q *QuickUnion) Find(a int) int {
	if a != q.arr[a] {
		q.arr[a] = q.Find(q.arr[a])
	}
	return q.arr[a]

}

func (q *QuickUnion) Union(a, b int) {
	pa, pb := q.Find(a), q.Find(b)
	if pa == pb {
		return
	}
	if q.size[pa] > q.size[pb] {
		q.arr[pb] = pa
	} else {
		q.arr[pa] = pb
	}
}

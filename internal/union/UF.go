package union

type UF interface {
	Connection(a, b int) bool
	Find(a int) int
	Union(a, b int)
	GetSize() int
}

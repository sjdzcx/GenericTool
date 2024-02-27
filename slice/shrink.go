package slice

// calCapacity 计算需要缩容的值和判断是否缩容
func calCapacity(c, l int) (int, bool) {
	//当长度小于 64 时，不需要缩容
	if c <= 64 {
		return c, false
	}
	if c <= 2048 && c/4 > l {
		return c / 2, true
	}

	if c > 2048 && c/2 > l {
		factor := 0.6
		return int(float32(c) * float32(factor)), true
	}
	return c, false

}

func Shrink[T any](src []T) []T {
	capacity, b := calCapacity(len(src), cap(src))
	if !b {
		return src
	}
	ret := make([]T, 0, capacity)
	ret = append(ret, src[:]...)
	return ret
}

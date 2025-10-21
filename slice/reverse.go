package slice

func Reverse[T any](s []T) []T {
	n := len(s)
	reversed := make([]T, n)
	for i, v := range s {
		reversed[n-i-1] = v
	}
	return reversed
}
func ReverseInPlace[T any](s []T) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}

}

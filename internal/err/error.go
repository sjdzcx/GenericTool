package err

import "fmt"

// NewErrIndexOutOfRange 创建一个代表下标超出范围的错误
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("slice: 下标超出范围，长度 %d, 下标 %d", length, index)
}

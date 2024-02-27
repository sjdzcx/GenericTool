package main

import (
	"GenericTool/slice"
	"fmt"
)

func main() {
	t := make([]int, 0, 3)
	t = append(t, 1, 3, 4)
	t, err := slice.Add[int](t, -1, -1)
	if err != nil {
		println(err.Error())

	} else {
		fmt.Printf("%+v", t)
	}

}

package main

import (
	"GenericTool/slice"
	"fmt"
)

type User struct {
	Name string
	age  int
}

func main() {
	t := make([]int, 0, 3)
	t = append(t, 1, 3, 4)
	t, err := slice.Add[int](t, -1, -1)
	if err != nil {
		println(err.Error())

	} else {
		fmt.Printf("%+v", t)
	}
	//ints, i, err := slice.Delete(t, 2)
	//println(i)

	/*	element, err := slice.DeleteElement(t, 3)
		if err != nil {
			println(err.Error())
		}

		fmt.Printf("%+v", element)*/

	//comparable()
	//测试结构体
	//u := make([]User, 0, 3)
	//u = append(u, User{Name: "hello3", age: 1})
	//u = append(u, User{Name: "hello2", age: 2})
	//u = append(u, User{Name: "hello1", age: 3})
	//element, err := slice.DeleteElement(u, User{Name: "hello2", age: 2})
	//fmt.Printf("%+v", element)

}

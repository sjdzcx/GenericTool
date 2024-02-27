package main

import (
	"fmt"
	"github.com/sjdzcx/GenericTool/union"
)

type User struct {
	Name string
	age  int
}

func test_union() {
	// 创建一个QuickUnion并查集实例
	qu := union.NewQuickUnion(10)

	// 进行一些连接操作
	qu.Union(1, 2)
	qu.Union(3, 4)
	qu.Union(5, 6)
	qu.Union(7, 8)
	qu.Union(9, 10)

	// 打印某些连接关系
	fmt.Println("Connection(1, 2):", qu.Connection(1, 2)) // 应为true
	fmt.Println("Connection(3, 5):", qu.Connection(3, 5)) // 应为false

	// 进行更多的连接操作
	qu.Union(2, 3)
	qu.Union(5, 7)

	// 再次打印某些连接关系
	fmt.Println("Connection(1, 5):", qu.Connection(1, 5)) // 应为true
	fmt.Println("Connection(2, 3):", qu.Connection(2, 3)) // 应为true

	println(qu.GetSize())
}
func main() {
	//t := make([]int, 0, 3)
	//t = append(t, 1, 3, 4)
	//t, err := slice.Add[int](t, -1, -1)
	//if err != nil {
	//	println(err.Error())
	//
	//} else {
	//	fmt.Printf("%+v", t)
	//}
	test_union()
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

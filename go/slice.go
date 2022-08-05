package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	// 简单切片
	s := a[1:3]
	fmt.Println(s)
	fmt.Println("len s:", len(s))
	fmt.Println("cap s:", cap(s))

	// 完整切片
	s1 := a[1:3:4]
	fmt.Println(s1)
	fmt.Println("len s1", len(s1))
	fmt.Println("cap s1", cap(s1))

	// make slice
	s2 := make([]bool, 2, 10)
	fmt.Println(s2)
	fmt.Println("len s2:", len(s2))
	fmt.Println("cap s2:", cap(s2))

	// 判断切片是否为空
	s3 := make([]bool, 0, 10)
	fmt.Println(s3)
	fmt.Println("len s3:", len(s3))
	fmt.Println("cap s3", cap(s3))

	// 未经过make初始化的切片不可以直接通过索引赋值，可以通过append函数追加元素。
	var s4 []string
	fmt.Println(s4 == nil)
	fmt.Println("len s4:", len(s4))
	fmt.Println("cap s4", cap(s4))
	fmt.Println(len(s4)) // 0

	s4 = append(s4, "test")
	fmt.Println(s4)
	fmt.Println("len s4:", len(s4))
	fmt.Println("cap s4", cap(s4))
	s4[0] = "test3"
	fmt.Println(s4)

	// 切片之间不能比较
	// slice can only be compared to nil
	s5 := []int{1, 2, 3, 4, 5}
	s6 := []int{1, 2, 3, 4, 5}
	fmt.Println(s5, s6)
	//fmt.Println(s5 == s6)

	// 切片的拷贝，浅copy，共享底层的数组
	s7 := make([]int, 4, 10)
	s8 := s7
	s7[0] = 100
	fmt.Println(s7)
	fmt.Println(s8)
}

package main

import (
	"fmt"
)

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
	// 判断slice是否为空的时候，用 len(slice) == 0来判断，不能用 slice == nil
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

	// 切片的浅copy，共享底层的数组
	s7 := make([]int, 4, 10)
	s8 := s7
	s7[0] = 100
	fmt.Println(s7)
	fmt.Println(s8)

	// 切片的遍历的两种方式：
	// 方式一：
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	// 方式二：利用for ... range遍历
	for k, v := range a {
		fmt.Println(k, v)
	}

	// append函数，用于扩展目标slice数据
	// 通过var声明的slice无需初始化，可以直接通过append添加item
	// 当底层数组容量不能满足append元素的时候，append按照一定的策略扩容数组。
	b := []int{1, 2, 3}
	b = append(b, 1, 2)
	fmt.Println(b)
	fmt.Println("len b", len(b))
	fmt.Println("cap b", cap(b))
	c := []int{5, 6}
	b = append(b, c...)
	fmt.Println(b)
	//d := [4]int{1, 3, 4}
	//b = append(b, d...)  不能append数组

	// 切片的扩容策略$GOROOT/src/runtime/slice.go
	// 首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
	// 否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
	// 否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
	// 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。

	// 切片中的深拷贝
	// copy的长度取决于dest的len大小
	e := []int{1, 2, 3, 4, 5}
	//var f []int 没有经过初始化copy不了
	f := make([]int, 4, 5)
	copy(f, e)
	fmt.Println("e:", e, len(e))
	fmt.Println("f:", f, len(f), cap(f))
	// 证明copy的长度取决于dest的len大小
	g := []int{1, 2}
	g = append(g, 1)
	fmt.Println(g, len(g), cap(g))
	copy(g, e)
	fmt.Println(g)

	// 从slice中删除元素
	h := []int{1, 2, 3, 4, 6}
	h = append(h[:2], h[3:]...)
	fmt.Println("h:", h)
}

package main

import "fmt"

func main() {
	// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
	// 定义map：map[KeyType]ValueType
	// map初始为nil，使用make分配内存
	var a map[int]string
	a = make(map[int]string, 10)

	// 未经初始化的map
	fmt.Println(a)
	fmt.Println("len a:", len(a))

	// 初始化map
	// 方式一
	a[1] = "test1"
	a[2] = "test2"
	fmt.Println(a)

	// 在声明时候填充元素，这样变量就不需要使用make初始化分配内存
	b := map[string]string{
		"name": "Tom",
		"age":  "13",
	}
	b["address"] = "shanghai"
	for i := range b {
		fmt.Println(i)
	}
	// 判断某值是否为map的key
	fmt.Println("--> 判断map中某key是否存在")
	value, ok := b["haha"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("查无此key~")
	}

	// map的遍历；
	// 1. 遍历key、value
	fmt.Println("--> 遍历map中的key, value")
	for k, v := range b {
		fmt.Println(k, v)
	}
	// 遍历key
	fmt.Println("--> 遍历map中的key元素")
	for k := range b {
		fmt.Println(k)
	}
	// 删除map中的元素
	fmt.Println("--> 删除map中的元素")
	delete(b, "xxx") // key不存在时不报错
	for k := range b {
		fmt.Println(k)
	}
	// 元素为map类型的切片
	fmt.Println("--> 元素为map类型的切片")
	mapSlice := make([]map[string]string, 3)
	mapSlice[0] = map[string]string{"s1": "s1"}

	for i, v := range mapSlice {
		fmt.Println(i, v)
	}
	// 值为切片类型的map
	fmt.Println("--> 值为切片类型的map")
	m1 := make(map[string][]string)
	m1["s1"] = []string{"s1", "s2", "s3"}
	for i, v := range m1 {
		fmt.Println(i, v)
		fmt.Printf("%T, %T", i, v)
	}
}

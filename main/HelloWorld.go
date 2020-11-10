package main

import "fmt"

func main() {

	fmt.Println(myFunc01())

	fmt.Println(myFunc02())

	a, b, c := myFunc03()
	fmt.Println(a, b, c)

	s1 := Student{1, 24, "kai", 'm', "shanghai"}
	fmt.Println(s1)

	var s2 Student

	s2.id = 1
	s2.name = "jay"
	s2.sex = 'm'
	//s2.add="dsad"

	fmt.Printf("%T", s1)

	fmt.Println(s2)

	s := Student{1, 25, "abc", 'x', "hangzhou"}

	test1(&s) //引用传递，类似浅拷贝
	test2(s)  //值传递,就是复制一个副本，各自改各自的

	fmt.Print(s.id)

}

func test1(s *Student) {
	s.id = 2
}

func test2(s Student) {
	s.id = 3
	println(s.id)
}

//func test1(p *Student) {
//	p.id=2
//	fmt.Print(p.id)
//
//
//}
//定义结构体
type Student struct {
	id   int
	age  int
	name string
	sex  byte
	add  string
}

//无参有返回值
func myFunc01() int {
	return 666
}

//给返回值命名，go推荐写法
//return空气就可以了
func myFunc02() (result int) {
	result = 6666
	return

}

//多个返回值
func myFunc03() (a int, b int, c int) {
	return 1, 2, 3

}

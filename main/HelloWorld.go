package main

import "fmt"


//无参有返回值
func myFunc01() int {
	return 666
}

//给返回值命名，go推荐写法
//return空气就可以了
func myFunc02() (result int) {
	result=6666
	return


}

//多个返回值
func myFunc03() (a int,b int,c int){
	return 1,2,3

}

func main()  {

	fmt.Println(myFunc01())

	fmt.Println(myFunc02())

	a,b,c:=myFunc03()
	fmt.Println(a,b,c)
}



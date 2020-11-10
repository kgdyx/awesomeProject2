package main

import (
	"encoding/json"
	"fmt"
)

//通过结构体生成json文本
//编码成json的话成员变量首字母必须大写
type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

func main() {
	//定义一个结构体变量 同时初始化
	s := IT{"bytedance", []string{"go", "python"}, true, 100.0}
	//编码成json,生成json文本

	buf, err := json.Marshal(s) //有两个返回值，一个是字节数组 切片，一个是错误,参数随便传一个就行
	if err != nil {
		fmt.Println("err=", err)
		return //有错误就别往下走了
	}
	fmt.Println("buf=", string(buf)) //没有错误就打印json文本，记得要把切片转成字符串数组

}

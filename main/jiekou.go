package main

import "fmt"

func main() {

	testa()
	testb(99)
	testc()

}

func testa() {
	fmt.Println("aaaaaaaaa")
}

func testb(x int) {

	var a [10]int
	a[x] = 111

	defer func() {

		if recover() != nil {
			fmt.Println(recover())
		}

	}()

}

func testc() {
	fmt.Println("cccccccc")
}

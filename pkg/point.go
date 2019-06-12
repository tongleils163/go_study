package pkg

import "fmt"


func Test() {
	// var s1 = []string{}
	// var s2 []string
	// fmt.Println(reflect.DeepEqual(s1, s2))
	// fmt.Println(s1)
	// fmt.Println(s2)
	// var m map[int]int
	// fmt.Println(m == nil)
	// fn(m)
	// fmt.Println(m == nil)
	//
	// var a []int
	// fmt.Println(a == nil)
	//
	// var c chan int
	// fmt.Println(c == nil)
	//
	// fmt.Println("fn1")
	// fn1(a)
	//
	// fn2(c)
}
func fn(m map[int]int) {
	m = make(map[int]int)
	fmt.Println(m == nil)
}

func fn1(a []int) {
	a = make([]int, 0)
	fmt.Println(a == nil)
}
func fn2(c chan int) {
	c = make(chan int)
	fmt.Println(c == nil)
}

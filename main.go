package main

import (
	"errors"
	"fmt"
)

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i <= 9 {
		s.data[s.i] = k
		s.i++
	}
}

func (s *stack) pop() int {
	if s.i > 0 {
		s.i--
		return s.data[s.i]
	}
	return 0
}

type Lesson struct {
	Id   int
	Name string
}
type LessonList struct {
	Lessons []*Lesson
}
type Struct1 struct {
	St1 Struct2
}

type Struct2 struct {
	Id int
}
type Animal interface {
	Bark()
}

type Dog struct {
}

func (d Dog) Bark() {
	fmt.Println("dog")
}

type Cat struct {
}

func (c *Cat) Bark() {
	fmt.Println("cat")
}

func Bark(a Animal) {
	a.Bark()
}

func getDog() Dog {
	return Dog{}
}

func getCat() Cat {
	return Cat{}
}

type S struct {
	A []S1
}
type S1 struct {
	I []int
}

func main() {
	s := &S{}
	s.A = append(s.A, S1{I: []int{5}})

	s.A[len(s.A)-1].I = append(s.A[len(s.A)-1].I, 6)

	fmt.Println(s)
	var a []int
	a = append(a, 4)
	// a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(len(a[3:5]))
	// index := []int{3, 7 }
	//
	// fmt.Println(a)
	// // for _, v := range index {
	// // 	a = append(a[:v], a[v+1:]...)
	// // }
	// for i := len(index) - 1; i >= 0; i-- {
	// 	a = append(a[:index[i]], a[index[i]+1:]...)
	// }
	// fmt.Println(a)

	// arr := []string{"a", "b", "c"}
	// fmt.Println(arr[0:2])
	// fmt.Println(arr[4 : 4+1])
	// lessons := &LessonList{Lessons: []*Lesson{{Id: 1, Name: "name1"}, {Id: 2, Name: "name2"}}}
	// // lesson1 := &Lesson{Id: 1, Name: "name1"}
	// var lesson2 *Lesson
	// lesson2 = lessons.Lessons[1]
	//
	// // fmt.Println(lessons.Lessons[1])
	// lesson2.Name = "name3"
	//
	// fmt.Println(lesson2)
	// // fmt.Println(lessons.Lessons[1])
	// // fmt.Println(lessons.Lessons[0])
	//
	// s := new(stack)
	//
	// s.push(5)
	// s.push(6)
	// fmt.Println(s)
	// fmt.Println(s.pop())
	// fmt.Println(s.pop())
	// // str :=" 你好 "
	// // fmt.Println(len([]byte(str)))
	// // fmt.Println(utf8.RuneCount([]byte(str)))
	// // pkg.TestChan2()
	// // for pos,_ := range  str{
	// // //	fmt.Println(pos)
	// // }
	// // fmt.Println(utf8.DecodeLastRune(str))
	// // s := "￿￿￿aA好"
	// r := [] rune(s)
	// copy (r[4:4+3], [] rune ("abc"))
	// fmt.Printf("Before: %s\n", s)
	// fmt.Printf("After : %s\n", string (r))

}

func RemoveSpringSlice(slice []string, start, end int) []string {
	return append(slice[:start], slice[end:]...)
}

const (
	Red   = 1
	Black = 2
)

type RBTreeNode struct {
	Color      int
	Value      int
	LeftChild  *RBTreeNode
	RightChild *RBTreeNode
	Parent     *RBTreeNode
}

type RBTree struct {
	Root *RBTreeNode
}

func (node *RBTreeNode) leftRotate() (*RBTreeNode, error) {
	var root *RBTreeNode
	if node == nil {
		return root, nil
	}
	if node.RightChild == nil {
		return root, errors.New("右子节点不可空")
	}
	parent := node.Parent
	var isLeft bool
	if parent != nil {
		isLeft = parent.LeftChild == node
	}
	grandson := node.RightChild.LeftChild

	node.RightChild.LeftChild = node
	node.Parent = node.RightChild
	node.RightChild = grandson

	// 是否根节点
	if parent == nil {
		node.Parent.Parent = nil
		root = node.Parent
	} else {
		if isLeft {
			parent.LeftChild = node.Parent
		} else {
			parent.RightChild = node.Parent
		}
		node.Parent.Parent = parent
	}
	return root, nil
}

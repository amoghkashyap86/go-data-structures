package main

import (
	"fmt"

	"github.com/amoghkashyap86/go-data-structures/ds"
)

func main() {
	some(1)
	some("amogh")

	s1 := new(ds.Stack[int])

	// s1.Push(3)
	// s1.Push(4)
	// isEmpty, ele := s1.Pop()

	// if isEmpty {
	// 	fmt.Println("stack is empty")
	// 	return
	// }

	// fmt.Println(ele)
	// isEmpty, ele = s1.Pop()

	for i := 0; i <= 1000; i++ {
		s1.Push(i)
	}

	isEmpty, _ := s1.Pop()

	if isEmpty {
		fmt.Println("stack is empty")
		return
	}

	for i := 0; i <= 500; i++ {
		s1.Pop()
	}

	// fmt.Println(len(s1.))

	// sl := []int{1, 2, 4}

	// fmt.Println(sl)
	// sl = sl[:2]
	// fmt.Println(sl)

}

func some[t any](param t) {
	fmt.Printf("type is %T anv value is %v\n", param, param)
}

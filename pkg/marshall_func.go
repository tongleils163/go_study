package pkg

import (
	"encoding/json"
	"fmt"
)

type Abc struct {
	A func(answer map[int]int) int
}

func (Abc) MarshallTestFunc(answer map[int]int) int {
	if answer[100] == 2 {
		return 5
	}
	return 6
}

func MarshallFunc() {
	var test Abc

	aa, err := json.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aa)
}

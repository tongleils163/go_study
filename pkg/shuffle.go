package pkg

import (
	"fmt"
	"math/rand"
	"sort"
)

func Shuffle()  {
	datas :=[]int{1,2,3,4,5,6,7,8,9,10}
	sort.Slice(datas, func(i, j int) bool {
		return rand.Float64() > 0.5
	})


	fmt.Println(datas)
}

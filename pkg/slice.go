package pkg

import (
	"fmt"
	"math/rand"
	"sort"
)

type str struct {
	Id   int
	Name string
}

func slice() {
	s := make([]*str, 0)
	s = append(s, &str{1, "1"})
	s = append(s, &str{2, "2"})
	s = append(s, &str{3, "3"})
	s = append(s, &str{4, "4"})
	for i := 0; i < 100; i++ {
		s = append(s, &str{i, "1"})
	}
	// for _, v := range s {
	// 	v.Id=5
	// }

	sort.Slice(s, func(i, j int) bool {
		s[i].Id = 2
		return rand.Float64() > 0.5
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i].Id < s[j].Id
	})
	// for _, v := range s {
	// 	fmt.Println(v.Id)
	// }

	fmt.Println(sort.SearchInts([]int{1, 2, 3}, 21))
}

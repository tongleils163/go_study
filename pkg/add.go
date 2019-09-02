package pkg

import "fmt"

type AN struct {
	Name string
}

func Add(a, b int) int {

	ans := make([]*AN, 0)
	ans = append(ans, &AN{Name: "AA"})

	a1 := *ans[0]
	fmt.Println(a1)
	a1.Name = "bb"
	fmt.Println()
	return a + b
}

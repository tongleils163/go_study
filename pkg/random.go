package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

func random() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 30; i++ {
		fmt.Println(r.Intn(3))
	}

	for i := 0; i < 20; i++ {
		fmt.Println(rand.Float64())
	}
}

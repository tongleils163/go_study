package pkg

import (
	"fmt"
	"time"
)

func TestGoFunc(data []*int) {

	go func() {

		time.Sleep(time.Second * 5)
		fmt.Println(data)
	}()

	time.Sleep(time.Second * 10)
}

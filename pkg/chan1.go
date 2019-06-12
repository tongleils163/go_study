package pkg

import (
	"fmt"
	"time"
)

func TestChan1()  {
	c := make(chan int)
	timeout := time.After(time.Second * 2)
	timer := time.NewTicker(time.Second * 3)
	var i int
	go func() {

		for {
			select {
			case <-c:
				fmt.Println("channel sign")
				return

			case <-timer.C:
				fmt.Println("3s")
			case <-timeout:
				i++
				fmt.Println(i, "2s")
			case <-time.After(time.Second * 4):
				fmt.Println("4s")
				fmt.Println("4s")
			default:
				fmt.Println("default")
				time.Sleep(time.Second)
			}

		}
	}()

	time.Sleep(time.Second * 8)
	close(c)
	fmt.Println("c close")
	time.Sleep(time.Second * 5)
}

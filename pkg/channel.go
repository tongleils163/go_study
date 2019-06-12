package pkg

import (
	"fmt"
	"time"
)

func ProducerCustomer() {

	data := make(chan int)
	go producer("生产者1", data)
	go producer("生产者2", data)
	go  consumer("消费者1", data)
	go  consumer("消费者2", data)

	// chan1 := make(chan int)
	// chan2 := make(chan int)
	//
	// go func() {
	// 	for i := 0; i < 2; i++ {
	// 		chan1 <- i
	// 	}
	// 	close(chan1)
	// }()
	//
	// go func() {
	// 	for x := range chan2 {
	// 		chan1 <- x
	// 	}
	// 	close(chan1)
	// }()
	//
	// for x := range chan1 {
	// 	fmt.Println(x)
	// }
}

func consumer(cname string, ch chan int) {
	for {
		select {
		case n := <-ch:
			fmt.Println(cname,n)
		case <-time.After(time.Second):
			return
		}
	}
	fmt.Println("ch closed.")
}

func producer(pname string, ch chan int) {
	for i := 0; i < 5; i++ {

		//fmt.Println("producer--", pname, ":", i)
		ch <- i
	}
}
